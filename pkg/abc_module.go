package gfoo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AbcModule struct {
	Module
}

func andImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	right := in.Pop()
	var rightOps []Op
	var err error
	
	if rightOps, err = right.Compile(in, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewAnd(form, rightOps)), nil
}

func branchImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var trueOps []Op
	var err error
	
	if trueOps, err = f.Compile(in, nil, scope); err != nil {
		return out, err
	}

	f = in.Pop()
	var falseOps []Op

	if falseOps, err = f.Compile(in, nil, scope); err != nil {
		return out, err
	}

	return append(out, NewBranch(form, trueOps, falseOps)), nil
}

func callImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	return append(out, NewCall(form, nil, nil)), nil
}

func callArgsImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	arg := in.Pop()
	var argOps []Op
	var err error
	
	if argOps, err = arg.Compile(in, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewCall(form, nil, argOps)), nil
}

func checkImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	cond := in.Pop()
	var condOps []Op
	var err error
	
	if condOps, err = cond.Compile(in, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewCheck(form, cond, condOps)), nil
}

func defineImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		Error(f.Pos(), "Expected id: %v", f)
	}
	
	var stack Slice
	
	if err := scope.EvalForm(in, &stack); err != nil {
		return out, err
	}
	
	val := stack.Pop()

	if val == nil {
	        return out, Error(id.Pos(), "Missing val: %v", id.name) 
	}
	

	if found := scope.Get(id.name); found == nil {
		scope.Set(id.name, *val)
	} else if found.scope != scope {
		return out, Error(id.Pos(), "Attempt to override compile time binding: %v", id.name)
	}

	return out, nil
}

func dropImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDrop(form)), nil
}

func dupImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDup(form)), nil
}

func forImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		return out, Error(form.Pos(), "Expected identifier: %v", f)
	}
	
	idIndex := -1
	
	if id.name != "_" {
		var err error
		
		if idIndex, err = scope.Let(id.name, form.Pos()); err != nil {
			return out, err
		}
	}

	body := in.Pop()	
	bodyOps, err := body.Compile(in, nil, scope)
	
	if err != nil {
		return out, err
	}

	return append(out, NewFor(form, idIndex, bodyOps)), nil
}

func includeImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	path, ok := f.(*Literal)
	
	if !ok {
		return out, Error(form.Pos(), "Invalid filename: %v", f)
	}

	if path.val.dataType != &TString {
		return out, Error(form.Pos(), "Invalid filename: %v", path.val)
	}

	return out, scope.Include(path.val.data.(string), func(forms []Form) error {
		for i := len(forms)-1; i >=0; i-- {
			in.Push(forms[i])
		}
		
		return nil
	})
}

func lambdaImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var args *Group
	var ok bool
	
	if args, ok = f.(*Group); !ok {
		return out, Error(form.Pos(), "Invalid argument list: %v", f)
	}

	var bodyOps []Op
	scope = NewScope(scope)

	for i := len(args.body)-1; i >= 0; i-- {
		a := args.body[i]
		var id *Id
		
		if id, ok = args.body[i].(*Id); !ok {
			return out, Error(a.Pos(), "Invalid argument: %v", a)
		}

		scope.Set(id.name, Undefined)
		index := len(scope.registers)
		scope.registers[id.name] = index
		bodyOps = append(bodyOps, NewLet(id, id.name, index))
	}

	body := in.Pop()
	var bodyForms []Form
	
	if s, ok := body.(*ScopeForm); ok {
		bodyForms = s.body
	} else {
		bodyForms = append(bodyForms, body)
	}
	
	var err error

	if bodyOps, err = scope.Compile(bodyForms, bodyOps); err != nil {
		return out, err
	}

	return append(out, NewLambdaOp(form, len(args.body), bodyOps)), nil
}

func letImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		Error(f.Pos(), "Expected id: %v", f)
	}

	index, err := scope.Let(id.name, form.Pos())

	if err != nil {
		return out, err
	}
	
	val := in.Pop()

	if out, err = val.Compile(in, out, scope); err != nil {
		return out, err
	}
			
	return append(out, NewLet(form, id.name, index)), nil
}

func macroImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		Error(f.Pos(), "Expected id: %v", f)
	}

	f = in.Pop()
	var args *Group

	if args, ok = f.(*Group); !ok {
		return out, Error(form.Pos(), "Invalid argument list: %v", f)
	}

	f = in.Pop()
	var body *ScopeForm

	if body, ok = f.(*ScopeForm); !ok {
		return out, Error(form.Pos(), "Invalid body: %v", f)
	}

	var bodyOps []Op
	macroScope := NewScope(scope)

	for i := len(args.body)-1; i >= 0 ; i-- {
		f := args.body[i]
		var id *Id
		
		if id, ok = f.(*Id); !ok {
			return out, Error(f.Pos(), "Invalid argument: %v", f)
		}
		
		macroScope.Set(id.name, Undefined)
		index := len(macroScope.registers)
		macroScope.registers[id.name] = index
		bodyOps = append(bodyOps, NewLet(f, id.name, index))
	}

	var err error
	
	if bodyOps, err = macroScope.Compile(body.body, bodyOps); err != nil {
		return out, err
	}
	
	scope.AddMacro(id.name, len(args.body), func(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
		var registers, stack Slice

		for i := 0; i < len(args.body); i++ {
			f := in.Pop()
			var v Val
			
			if v, err = f.Quote(in, scope, nil, &registers, f.Pos()); err != nil {
				return out, err
			}
			
			stack.Push(v)
		}

		if err := EvalOps(bodyOps, nil, &registers, &stack); err != nil {
			return out, err
		}

		scope = NewScope(nil)
		
		for i := stack.Len()-1; i >= 0; i-- {
			in.Push(stack.items[i].Unquote(scope, form.Pos()))
		}

		return out, nil
	})
		
	return out, nil
}

func mapImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		return out, Error(form.Pos(), "Expected identifier: %v", f)
	}
	
	idIndex := -1
	
	if id.name != "_" {
		var err error
		
		if idIndex, err = scope.Let(id.name, form.Pos()); err != nil {
			return out, err
		}
	}

	body := in.Pop()	
	bodyOps, err := body.Compile(in, nil, scope)
	
	if err != nil {
		return out, err
	}
	
	return append(out, NewMap(form, idIndex, bodyOps)), nil
}

func methodImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		Error(f.Pos(), "Expected id: %v", f)
	}

	f = in.Pop()
	var argsForm *Group

	if argsForm, ok = f.(*Group); !ok || len(argsForm.body) < 1 {
		return out, Error(form.Pos(), "Invalid argument list: %v", f)
	}

	var retsForm *Group

	if retsForm, ok = argsForm.body[len(argsForm.body)-1].(*Group); !ok {
		return out, Error(f.Pos(), "Invalid result list: %v", f)
	}

	body := in.Pop()
	var args []Arg
	
	for i := 0; i < len(argsForm.body)-1; i++ {
		anf := argsForm.body[i]
		var ans []string
		
		if id, ok := anf.(*Id); ok {
			ans = append(ans, id.name)
		} else if g, ok := anf.(*Group); ok {
			for _, anf = range g.body {
				if id, ok := anf.(*Id); ok {
					ans = append(ans, id.name)
				} else {
					return out, Error(anf.Pos(), "Invalid argument name: %v", anf)
				}
			}
		} else {
			return out, Error(anf.Pos(), "Invalid argument name: %v", anf)
		}
		
		i++
		atf := argsForm.body[i]

		if ati, ok := atf.(*Literal); ok && ati.val.dataType == &TInt {
			for _, an := range ans {
				args = append(args, AIndex(an, int(ati.val.data.(Int))))
			}
		} else if atn, ok := atf.(*Id); ok {
			atb := scope.Get(atn.name)

			if atb == nil {
				return out, Error(atf.Pos(), "Type not found: %v", atn)
			}

			if atb.val.dataType != &TMeta {
				return out, Error(atf.Pos(), "Expected type: %v", atb.val.dataType.Name())
			}

			for _, an := range ans {
				args = append(args, AType(an, atb.val.data.(Type)))
			}
		} else {
			return out, Error(atf.Pos(), "Invalid argument type: %v", atf)
		}
	}

	var rets []Ret

	for _, f := range retsForm.body {
		if rti, ok := f.(*Literal); ok && rti.val.dataType == &TInt {
			rets = append(rets, RIndex(int(rti.val.data.(Int))))
		} else if rtn, ok := f.(*Id); ok {
			rtb := scope.Get(rtn.name)
			
			if rtb == nil {
				return out, Error(f.Pos(), "Type not found: %v", rtn)
			}

			if rtb.val.dataType != &TMeta {
				return out, Error(f.Pos(), "Expected type: %v", rtb.val.dataType.Name())
			}

			rets = append(rets, RType(rtb.val.data.(Type)))
		}
	}
	
	var bodyOps []Op
	methodScope := NewScope(scope)
	
	for i := len(args)-1; i >= 0; i-- {
		a := args[i]
		
		if a.name != "_" {
			methodScope.Set(a.name, Undefined)
			index := len(methodScope.registers)
			methodScope.registers[a.name] = index
			bodyOps = append(bodyOps, NewLet(argsForm, a.name, index))
		}
	}

	var err error
	var bodyForms []Form
	
	if s, ok := body.(*ScopeForm); ok {
		bodyForms = s.body
	} else {
		bodyForms = append(bodyForms, body)
	}
	
	if bodyOps, err = methodScope.Compile(bodyForms, bodyOps); err != nil {
		return out, err
	}
	
	m := scope.AddMethod(id.name, args, rets, func(thread *Thread, registers, stack *Slice, pos Pos) error {
		return EvalOps(bodyOps, thread, registers, stack)
	})

	return append(out, NewMethodOp(form, m)), nil
}

func orImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	right := in.Pop()
	var rightOps []Op
	var err error
	
	if rightOps, err = right.Compile(in, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewOr(form, rightOps)), nil
}

func pauseImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	result := in.Pop()
	var resultOps []Op
	var err error
	
	if resultOps, err = result.Compile(in, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewPause(form, resultOps)), nil
}

func threadImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var args *Group
	var ok bool
	
	if args, ok = f.(*Group); !ok {
		return out, Error(form.Pos(), "Invalid argument list: %v", f)
	}

	var argOps []Op
	var err error
	
	if argOps, err = args.Compile(in, nil, scope); err != nil {
		return out, err
	}

	f = in.Pop()
	var body *ScopeForm
	
	if body, ok = f.(*ScopeForm); !ok {
		return out, Error(form.Pos(), "Invalid body: %v", f)
	}
	
	var bodyOps []Op
	
	if bodyOps, err = scope.Compile(body.body, nil); err != nil {
		return out, err
	}
	
	return append(out, NewThreadOp(form, argOps, bodyOps)), nil
}

func timesImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	body := in.Pop()	
	bodyOps, err := body.Compile(in, nil, scope)
	
	if err != nil {
		return out, err
	}
	
	return append(out, NewTimes(form, bodyOps)), nil
}

func traitImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	var f Form
	
	if f = in.Pop(); f == nil {
		return out, Error(form.Pos(), "Missing id")
	}

	id, ok := f.(*Id)
	
	if !ok {
		return out, Error(form.Pos(), "Expected id: %v", f)
	}

	var stack Slice

	if err := scope.EvalForm(in, &stack); err != nil {
		return out, err
	}

	var parents []Type
		
	for _, v := range stack.items {
		if v.dataType != &TMeta {
			return out, Error(form.Pos(), "Expected type: %v", v)
		}

		pt, ok := v.data.(*Trait);
		
		if !ok {
			return out, Error(form.Pos(), "Expected trait: %v", v)
		}

		parents = append(parents, pt)
	}

	scope.AddType(NewTrait(id.name, parents...))
	return out, nil
}

func typeImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	id, ok := f.(*Id)
	
	if !ok {
		return out, Error(form.Pos(), "Expected id: %v", f)
	}

	var stack Slice

	if err := scope.EvalForm(in, &stack); err != nil {
		return out, err
	}

	var parents []Type
		
	for _, v := range stack.items {
		if v.dataType != &TMeta {
			return out, Error(form.Pos(), "Expected type: %v", v)
		}

		pt, ok := v.data.(*Trait);
		
		if !ok {
			return out, Error(form.Pos(), "Expected trait: %v", v)
		}

		parents = append(parents, pt)
	}

	stack.Clear()
	
	if err := scope.EvalForm(in, &stack); err != nil {
		return out, err
	}
	
	imp := stack.Pop()

	if imp == nil {
		return out, Error(form.Pos(), "Type not found")
	}

	if imp.dataType != &TMeta {
		return out, Error(form.Pos(), "Expected type: %v", imp)
	}

	impType, ok := imp.data.(ValType)

	if !ok {
		return out, Error(form.Pos(), "Expected value type: %v", imp)
	}

	parents = append(parents, impType)
	t := impType.New(id.name, parents...)
	scope.AddType(t)

	scope.AddMethod(fmt.Sprintf("as-%v", strings.ToLower(t.Name())),
		[]Arg{AType("val", impType)},
		[]Ret{RType(t)},
		func(thread *Thread, registers, stack *Slice, pos Pos) error {
			v := stack.Peek()
			v.dataType = t
			return nil
		})
	
	return out, nil
}

func useImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	sourceForm := in.Pop()
	id, ok := sourceForm.(*Id)
	
	if !ok {
		return out, Error(form.Pos(), "Expected identifier: %v", sourceForm)
	}

	source := scope.Get(id.name)

	if source == nil || source.val == Undefined {
		return out, Error(form.Pos(), "Unknown identifier: %v", id)
	}

	namesForm := in.Pop()
	var names []string

	if f, ok := namesForm.(*Group); ok {
		for _, f := range f.body {
			id, ok = f.(*Id)
			
			if !ok {
				return out, Error(f.Pos(), "Expected identifier: %v", f)
			}

			names = append(names, id.name)
		}
	} else if id, ok := namesForm.(*Id); !ok || id.name != "..." {
		return out, Error(form.Pos(), "Invalid identifier list: %v", namesForm)
	}
	
	return out, scope.Use(source.val, names, form.Pos())
}

func breakImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	return &Break
}

func cloneImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(stack.Pop().Clone())
	return nil
}

func dumpImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Pop().Dump(os.Stdout)
	os.Stdout.WriteString("\n")
	return nil
}

func eqImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Eq))
	return nil
}

func gtImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Gt))
	return nil
}

func gteImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) >= Eq))
	return nil
}

func intAddImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, stack.Pop().data.(Int) + stack.Pop().data.(Int)))
	return nil
}

func intCountImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in, err := stack.Pop().Iter(pos)

	if err != nil {
		return err
	}

	stack.Push(NewVal(&TIter, in))
	return nil
}

func intDecImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	v := stack.Pop().data.(Int)
	stack.Push(NewVal(&TInt, v-1))
	return nil
}

func intIncImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	v := stack.Pop().data.(Int)
	stack.Push(NewVal(&TInt, v+1))
	return nil
}

func intMulImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, stack.Pop().data.(Int) * stack.Pop().data.(Int)))
	return nil
}

func intSubImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	y := stack.Pop().data.(Int)
	stack.Push(NewVal(&TInt, stack.Pop().data.(Int) - y))
	return nil
}

func intToStringImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TString, strconv.FormatInt(stack.Pop().data.(Int), 10)))
	return nil
}

func isImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TBool, stack.Pop().Is(*stack.Pop())))
	return nil
}

func isaImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	parent := stack.Pop().data.(Type)
	out := stack.Pop().data.(Type).Isa(parent)
	
	if out == nil {
		stack.Push(Nil)
	} else {
		stack.Push(NewVal(&TMeta, out))
	}
	
	return nil
}

func ltImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Lt))
	return nil
}

func lteImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) <= Eq))
	return nil
}

func sayImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Pop().Print(os.Stdout)
	os.Stdout.WriteString("\n")
	return nil
}

func sliceItemsImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in, err := stack.Pop().Iter(pos)

	if err != nil {
		return err
	}

	stack.Push(NewVal(&TIter, in))
	return nil
}

func sliceLengthImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, Int(stack.Pop().data.(*Slice).Len())))
	return nil
}

func slicePeekImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in := stack.Pop().data.(*Slice).Peek()
	var out Val
	
	if in == nil {
		out = Nil
	} else {
		out = *in
	}
	
	stack.Push(out)
	return nil
}

func slicePopImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in := stack.Pop().data.(*Slice).Pop()
	var out Val
	
	if in == nil {
		out = Nil
	} else {
		out = *in
	}
	
	stack.Push(out)
	return nil
}

func slicePushImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	v := stack.Pop()
	stack.Pop().data.(*Slice).Push(*v)
	return nil
}

func spreadImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in, err := stack.Pop().Iter(pos)

	if err != nil {
		return err
	}
	
	return in.For(func(v Val, thread *Thread, pos Pos) error {
		stack.Push(v)
		return nil
	}, thread, pos)
}

func stringCharsImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in, err := stack.Pop().Iter(pos)

	if err != nil {
		return err
	}

	stack.Push(NewVal(&TIter, in))
	return nil
}

func stringLengthImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, Int(len(stack.Pop().data.(string)))))
	return nil
}

func threadWaitImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Pop().data.(*Thread).Wait(stack, pos)
	return nil
}

func toBoolImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TBool, stack.Pop().Bool()))
	return nil
}

func typeofImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TMeta, stack.Pop().dataType))
	return nil
}

func (self *AbcModule) Init() *Module {
	self.Module.Init()
	
	self.AddType(&TAny)
	self.AddType(&TBool)
	self.AddType(&TCall)
	self.AddType(&TChar)
	self.AddType(&TFunction)
	self.AddType(&TId)
	self.AddType(&TInt)
	self.AddType(&TLambda)
	self.AddType(&TMacro)
	self.AddType(&TMeta)
	self.AddType(&TMethod)
	self.AddType(&TModule)
	self.AddType(&TNil)
	self.AddType(&TNumber)
	self.AddType(&TOption)
	self.AddType(&TPair)
	self.AddType(&TScope)
	self.AddType(&TSequence)
	self.AddType(&TSlice)
	self.AddType(&TString)

	self.AddVal("NIL", &TNil, nil)
	self.AddVal("T", &TBool, true)
	self.AddVal("F", &TBool, false)

	self.AddMacro("and:", 1, andImp)
	self.AddMacro("?:", 2, branchImp)
 	self.AddMacro("call", 0, callImp)
 	self.AddMacro("call:", 1, callArgsImp)
	self.AddMacro("check:", 1, checkImp)
	self.AddMacro("define:", 2, defineImp)
	self.AddMacro("_", 0, dropImp)
	self.AddMacro("..", 0, dupImp)
 	self.AddMacro("for:", 2, forImp)
	self.AddMacro("include:", 1, includeImp)
	self.AddMacro("/:", 2, lambdaImp)
	self.AddMacro("let:", 2, letImp)
	self.AddMacro("macro:", 3, macroImp)
 	self.AddMacro("map:", 2, mapImp)
	self.AddMacro("method:", 3, methodImp)
	self.AddMacro("or:", 1, orImp)
	self.AddMacro("pause:", 1, pauseImp)
	self.AddMacro("thread:", 2, threadImp)
 	self.AddMacro("times:", 1, timesImp)
	self.AddMacro("trait:", 2, traitImp)
	self.AddMacro("type:", 2, typeImp)
	self.AddMacro("use:", 2, useImp)

	self.AddMethod("break", nil, nil, breakImp)
	self.AddMethod("clone", []Arg{AType("val", &TAny)}, []Ret{RIndex(0)}, cloneImp)
	self.AddMethod("dump", []Arg{AType("val", &TOption)}, nil, dumpImp)
	self.AddMethod("=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, eqImp)
	self.AddMethod(">", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, gtImp)
	self.AddMethod(">=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, gteImp)
	self.AddMethod("+", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intAddImp)
	self.AddMethod("count", []Arg{AType("val", &TInt)}, []Ret{RType(&TIter)}, intCountImp)
	self.AddMethod("-1", []Arg{AType("val", &TInt)}, []Ret{RType(&TInt)}, intDecImp)
	self.AddMethod("+1", []Arg{AType("val", &TInt)}, []Ret{RType(&TInt)}, intIncImp)
	self.AddMethod("*", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intMulImp)
	self.AddMethod("-", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intSubImp)

	self.AddMethod("to-string", []Arg{AType("val", &TInt)}, []Ret{RType(&TString)}, intToStringImp)

	self.AddMethod("is", []Arg{AType("x", &TOption), AType("y", &TOption)}, []Ret{RType(&TBool)}, isImp)
	
	self.AddMethod("isa",
		[]Arg{AType("child", &TMeta), AType("parent", &TMeta)},
		[]Ret{RType(Option(&TMeta))},
		isaImp)
	
	self.AddMethod("<", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, ltImp)
	self.AddMethod("<=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, lteImp)
	self.AddMethod("say", []Arg{AType("val", &TAny)}, nil, sayImp)
	self.AddMethod("items", []Arg{AType("val", &TSlice)}, []Ret{RType(&TIter)}, sliceItemsImp)
	self.AddMethod("length", []Arg{AType("val", &TSlice)}, []Ret{RType(&TInt)}, sliceLengthImp)
	self.AddMethod("peek", []Arg{AType("val", &TSlice)}, []Ret{RType(&TOption)}, slicePeekImp)
	self.AddMethod("pop", []Arg{AType("val", &TSlice)}, []Ret{RType(&TOption)}, slicePopImp)
	self.AddMethod("push", []Arg{AType("target", &TSlice), AType("val", &TAny)}, nil, slicePushImp)
	self.AddMethod("...", []Arg{AType("val", &TSequence)}, nil, spreadImp)
	self.AddMethod("chars", []Arg{AType("val", &TString)}, []Ret{RType(&TIter)}, stringCharsImp)
	self.AddMethod("length", []Arg{AType("val", &TString)}, []Ret{RType(&TInt)}, stringLengthImp)
	self.AddMethod("wait", []Arg{AType("thread", &TThread)}, nil, threadWaitImp)
	self.AddMethod("to-bool", []Arg{AType("val", &TAny)}, []Ret{RType(&TBool)}, toBoolImp)
	self.AddMethod("typeof", []Arg{AType("val", &TAny)}, []Ret{RType(&TMeta)}, typeofImp)

	self.Eval(`
macro: all: (body) {
  '{
     let: #in ()
     T #in for: #v and: (#v @body)
   }
}

macro: any: (body) {
  '{
     let: #in ()
     F #in for: #v (#v @body if: (or: T break))
   }
}

macro: else: (body) {
  '(?: () @body)
}

macro: if: (body) {
  '(?: @body ())
}
        `, nil, nil, nil)
	
	return &self.Module
}
