package gfoo

import (
	"os"
)

func branchImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	var trueOps []Op
	var err error
	
	if trueOps, err = f.Compile(nil, nil, scope); err != nil {
		return out, err
	}

	f = in.Pop()
	var falseOps []Op

	if falseOps, err = f.Compile(nil, nil, scope); err != nil {
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
	
	if argOps, err = arg.Compile(nil, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewCall(form, nil, argOps)), nil
}

func checkImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	cond := in.Pop()
	var condOps []Op
	var err error
	
	if condOps, err = cond.Compile(nil, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewCheck(form, cond, condOps)), nil
}

func defineImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		scope.Error(f.Pos(), "Expected id: %v", f)
	}
	
	var stack Slice
	
	if err := scope.EvalForm(in, &stack); err != nil {
		return out, err
	}
	
	val := stack.Pop()

	if val == nil {
	        return out, scope.Error(id.Pos(), "Missing val: %v", id.name) 
	}
	

	if found := scope.Get(id.name); found == nil {
		scope.Set(id.name, *val)
	} else if found.scope != scope {
		return out, scope.Error(id.Pos(), "Attempt to override compile time binding: %v", id.name)
	}

	return out, nil
}

func doImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	body, ok := f.(*ScopeForm)

	if !ok {
		return out, scope.Error(form.Pos(), "Invalid body: %v", f)
	}

	bodyOps, err := scope.Clone().Compile(body.body, nil)
	
	if err != nil {
		return out, err
	}
	
	return append(out, NewScopeOp(form, bodyOps, nil)), nil
}

func dropImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDrop(form)), nil
}

func dupImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDup(form)), nil
}

func includeImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	path, ok := f.(*Literal)
	
	if !ok {
		return out, scope.Error(form.Pos(), "Invalid filename: %v", f)
	}

	if path.val.dataType != &TString {
		return out, scope.Error(form.Pos(), "Invalid filename: %v", path.val)
	}

	return out, scope.Include(path.val.data.(string), func(forms []Form) error {
		for i := len(forms)-1; i >=0; i-- {
			in.Push(forms[i])
		}
		
		return nil
	})
}

func isImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewIs(form, nil)), nil
}

func lambdaImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var args *Group
	var ok bool
	
	if args, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid argument list: %v", f)
	}

	f = in.Pop()
	var body *ScopeForm

	if body, ok = f.(*ScopeForm); !ok {
		return out, scope.Error(form.Pos(), "Invalid body: %v", f)
	}
	
	var bodyOps []Op

	for i := len(args.body)-1; i >= 0; i-- {
		a := args.body[i]
		var id *Id
		
		if id, ok = args.body[i].(*Id); !ok {
			return out, scope.Error(a.Pos(), "Invalid argument: %v", a)
		}
		
		bodyOps = append(bodyOps, NewLet(id, id.name))
	}
	
	scope = scope.Clone()
	var err error
	
	if bodyOps, err = scope.Compile(body.body, bodyOps); err != nil {
		return out, err
	}

	return append(out, NewPush(form, NewVal(&TLambda, NewLambda(len(args.body), bodyOps, scope)))), nil
}

func letImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		scope.Error(f.Pos(), "Expected id: %v", f)
	}

	if found := scope.Get(id.name); found == nil {
		scope.Set(id.name, Nil)
	} else if found.val != Nil {
	        return out, scope.Error(id.Pos(), "Attempt to override compile time binding: %v", id.name)
	} else if found.scope != scope {
		found.Init(scope, Nil)
	} else {
	        return out, scope.Error(id.Pos(), "Duplicate binding: %v", id.name) 
	}
	
	val := in.Pop()
	var err error

	if out, err = val.Compile(in, out, scope); err != nil {
		return out, err
	}
			
	return append(out, NewLet(form, id.name)), nil
}

func macroImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		scope.Error(f.Pos(), "Expected id: %v", f)
	}

	f = in.Pop()
	var args *Group

	if args, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid argument list: %v", f)
	}

	f = in.Pop()
	var body *ScopeForm

	if body, ok = f.(*ScopeForm); !ok {
		return out, scope.Error(form.Pos(), "Invalid body: %v", f)
	}

	var bodyOps []Op

	for i := len(args.body)-1; i >= 0; i-- {
		f := args.body[i]
		var id *Id
		
		if id, ok = f.(*Id); !ok {
			return out, scope.Error(f.Pos(), "Invalid argument: %v", f)
		}
		
		bodyOps = append(bodyOps, NewLet(f, id.name))
	}

	var err error
	macroScope := scope.Clone()
	
	if bodyOps, err = macroScope.Compile(body.body, bodyOps); err != nil {
		return out, err
	}
	
	scope.AddMacro(id.name, len(args.body), func(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
		var stack Slice

		for i := 0; i < len(args.body); i++ {
			f := in.Pop()
			var v Val
			
			if v, err = f.Quote(scope, f.Pos()); err != nil {
				return out, err
			}
			
			stack.Push(v)
		}

		scope = macroScope.Clone()
		
		if err := scope.EvalOps(bodyOps, &stack); err != nil {
			return out, err
		}

		for i := stack.Len()-1; i >= 0; i-- {
			in.Push(stack.items[i].Unquote(scope, form.Pos()))
		}

		return out, nil
	})
		
	return out, nil
}

func methodImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	id, ok := f.(*Id)

	if !ok {
		scope.Error(f.Pos(), "Expected id: %v", f)
	}

	f = in.Pop()
	var argsForm *Group

	if argsForm, ok = f.(*Group); !ok || len(argsForm.body) < 1 {
		return out, scope.Error(form.Pos(), "Invalid argument list: %v", f)
	}

	var retsForm *Group

	if retsForm, ok = argsForm.body[len(argsForm.body)-1].(*Group); !ok {
		return out, scope.Error(f.Pos(), "Invalid result list: %v", f)
	}

	body := in.Pop()
	var args []Arg
	
	for i := 0; i < len(argsForm.body)-1; i++ {
		anf := argsForm.body[i]
		an := anf.(*Id).name
		i++
		atf := argsForm.body[i]

		if ati, ok := atf.(*Literal); ok && ati.val.dataType == &TInt {
			args = append(args, AIndex(an, int(ati.val.data.(*Int).Int64())))
		} else if atn, ok := atf.(*Id); ok {
			atb := scope.Get(atn.name)

			if atb == nil {
				return out, scope.Error(atf.Pos(), "Type not found: %v", atn)
			}

			if atb.val.dataType != &TMeta {
				return out, scope.Error(atf.Pos(), "Expected type: %v", atb.val.dataType.Name())
			}

			args = append(args, AType(an, atb.val.data.(Type)))
		} else {
			return out, scope.Error(atf.Pos(), "Invalid argument type: %v", atf)
		}
	}

	var rets []Ret

	for _, f := range retsForm.body {
		if rti, ok := f.(*Literal); ok && rti.val.dataType == &TInt {
			rets = append(rets, RIndex(int(rti.val.data.(*Int).Int64())))
		} else if rtn, ok := f.(*Id); ok {
			rtb := scope.Get(rtn.name)
			
			if rtb == nil {
				return out, scope.Error(f.Pos(), "Type not found: %v", rtn)
			}

			if rtb.val.dataType != &TMeta {
				return out, scope.Error(f.Pos(), "Expected type: %v", rtb.val.dataType.Name())
			}

			rets = append(rets, RType(rtb.val.data.(Type)))
		}
	}
	
	var bodyOps []Op
	
	for i := len(args)-1; i >= 0; i-- {
		if args[i].name != "_" {
			bodyOps = append(bodyOps, NewLet(argsForm, args[i].name))
		}
	}

	methodScope := scope.Clone()
	var err error
	var bodyForms []Form
	var cloneScope bool
	
	if s, ok := body.(*ScopeForm); ok {
		bodyForms = s.body
		cloneScope = true
	} else {
		bodyForms = append(bodyForms, body)
		cloneScope = false
	}
	
	if bodyOps, err = methodScope.Compile(bodyForms, bodyOps); err != nil {
		return out, err
	}
	
	scope.AddMethod(id.name, args, rets, func(scope *Scope, stack *Slice, pos Pos) error {
		if cloneScope {
			methodScope = methodScope.Clone()
		}
		
		return methodScope.EvalOps(bodyOps, stack)
	})
	
	return out, nil
}

func pairImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewPairOp(form)), nil
}

func pauseImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	result := in.Pop()
	var resultOps []Op
	var err error
	
	if resultOps, err = result.Compile(nil, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewPause(form, resultOps)), nil
}

func scopeImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var bindings *Group
	var ok bool

	if bindings, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid bindings: %v", f)
	}

	bindingForms := NewForms(bindings.body)
	var keys []string
	var values []Op

	for {
		if f = bindingForms.Pop(); f == nil {
			break
		}

		id, ok := f.(*Id)
		
		if !ok {
			return out, scope.Error(f.Pos(), "Expected id: %v", f)
		}

		keys = append(keys, id.name)
		
		if f = bindingForms.Pop(); f == nil {
			return out, scope.Error(id.Pos(), "Missing value: %v", id)
		}

		var err error
		
		if values, err = f.Compile(bindingForms, values, scope); err != nil {
			return out, err
		}
	}

	return append(out, NewScopeDef(form, keys, values)), nil
}

func threadImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var args *Group
	var ok bool
	
	if args, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid argument list: %v", f)
	}

	var argOps []Op
	var err error
	
	if argOps, err = args.Compile(nil, nil, scope); err != nil {
		return out, err
	}

	f = in.Pop()
	var body *ScopeForm
	
	if body, ok = f.(*ScopeForm); !ok {
		return out, scope.Error(form.Pos(), "Invalid body: %v", f)
	}
	
	var bodyOps []Op
	
	if bodyOps, err = scope.Compile(body.body, nil); err != nil {
		return out, err
	}
	
	return append(out, NewThreadOp(form, argOps, bodyOps)), nil
}

func typeDefImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	traitId, ok := f.(*Id)
	
	if !ok {
		return out, scope.Error(form.Pos(), "Expected id: %v", f)
	}

	var stack Slice

	if err := scope.EvalForm(in, &stack); err != nil {
		return out, err
	}
	
	imp := stack.Pop()

	if imp == nil {
		return out, scope.Error(form.Pos(), "Type not found")
	}

	if imp.dataType != &TMeta {
		return out, scope.Error(form.Pos(), "Expected type: %v", imp)
	}

	impType, ok := imp.data.(ValType)

	if !ok {
		return out, scope.Error(form.Pos(), "Expected value type: %v", imp)
	}
	
	t := impType.New(traitId.name, impType)
	scope.AddType(t)

	scope.AddMethod("as",
		[]Arg{AType("val", impType), AVal("type", NewVal(&TMeta, t))},
		[]Ret{RType(t)},
		func(scope *Scope, stack *Slice, pos Pos) error {
			stack.Pop()
			v := stack.Peek()
			v.dataType = t
			return nil
		})
	
	return out, nil
}

func useImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	f := in.Pop()
	sourceForm, ok := f.(*Id)
	
	if !ok {
		return out, scope.Error(form.Pos(), "Invalid source: %v", f)
	}

	source := scope.Get(sourceForm.name);

	if source == nil || source.val == Nil {
		return out, scope.Error(form.Pos(), "Unknown identifier: %v", f)
	}
	
	f = in.Pop()
	ids, ok := f.(*Group)
	
	if !ok {
		return out, scope.Error(form.Pos(), "Invalid identifier list: %v", f)
	}

	for _, f := range ids.body {
		k, ok := f.(*Id)
		
		if !ok {
			return out, scope.Error(f.Pos(), "Invalid identifier: %v", f)
		}

		v, err := source.val.Get(k.name, scope, k.Pos())

		if err != nil {
			return out, err
		}

		if scope.Get(k.name) != nil {
			return out, scope.Error(k.Pos(), "Duplicate identifier: %v", k)
		}

		scope.Set(k.name, v)
	}

	return out, nil
}

func boolImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TBool, stack.Pop().Bool()))
	return nil
}

func cloneImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(stack.Pop().Clone())
	return nil
}

func dumpImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Pop().Dump(os.Stdout)
	os.Stdout.WriteString("\n")
	return nil
}

func eqImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Eq))
	return nil
}

func gtImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Gt))
	return nil
}

func gteImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) >= Eq))
	return nil
}

func intAddImp(scope *Scope, stack *Slice, pos Pos) error {
	var z Int
	z.Add(stack.Pop().data.(*Int), stack.Pop().data.(*Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func intMulImp(scope *Scope, stack *Slice, pos Pos) error {
	var z Int
	z.Mul(stack.Pop().data.(*Int), stack.Pop().data.(*Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func intSpreadImp(scope *Scope, stack *Slice, pos Pos) error {
	v := stack.Pop().data.(*Int)
	
	for i := int64(0); i < v.Int64(); i++  {
		stack.Push(NewVal(&TInt, NewInt(i)))
	}

	return nil
}

func intSubImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop()
	var z Int
	z.Sub(stack.Pop().data.(*Int), y.data.(*Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func isaImp(scope *Scope, stack *Slice, pos Pos) error {
	parent := stack.Pop().data.(Type)
	out := stack.Pop().data.(Type).Isa(parent)
	
	if out == nil {
		stack.Push(Nil)
	} else {
		stack.Push(NewVal(&TMeta, out))
	}
	
	return nil
}

func loadImp(scope *Scope, stack *Slice, pos Pos) error {
	return scope.Load(stack.Pop().data.(string), stack)
}

func ltImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Lt))
	return nil
}

func lteImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) <= Eq))
	return nil
}

func pairSpreadImp(scope *Scope, stack *Slice, pos Pos) error {
	p := stack.Pop().data.(Pair)
	stack.Push(p.left)
	stack.Push(p.right)
	return nil
}

func sayImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Pop().Print(os.Stdout)
	os.Stdout.WriteString("\n")
	return nil
}

func sliceLengthImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, NewInt(int64(stack.Pop().data.(*Slice).Len()))))
	return nil
}

func slicePeekImp(scope *Scope, stack *Slice, pos Pos) error {
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

func slicePopImp(scope *Scope, stack *Slice, pos Pos) error {
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

func slicePushImp(scope *Scope, stack *Slice, pos Pos) error {
	v := stack.Pop()
	stack.Pop().data.(*Slice).Push(*v)
	return nil
}

func sliceSpreadImp(scope *Scope, stack *Slice, pos Pos) error {
	for _, v := range stack.Pop().data.(*Slice).items {
		stack.Push(v)
	}

	return nil
}

func stringLengthImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, NewInt(int64(len(stack.Pop().data.(string))))))
	return nil
}

func stringSpreadImp(scope *Scope, stack *Slice, pos Pos) error {
	for _, c := range stack.Pop().data.(string) {
		stack.Push(NewVal(&TChar, c))
	}

	return nil
}

func typeImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TMeta, stack.Pop().dataType))
	return nil
}

func (self *Scope) InitAbc() *Scope {
	self.AddType(&TAny)
	self.AddType(&TBool)
	self.AddType(&TChar)
	self.AddType(&TFunction)
	self.AddType(&TId)
	self.AddType(&TInt)
	self.AddType(&TLambda)
	self.AddType(&TMacro)
	self.AddType(&TMeta)
	self.AddType(&TMethod)
	self.AddType(&TNil)
	self.AddType(&TNumber)
	self.AddType(&TOptional)
	self.AddType(&TPair)
	self.AddType(&TScope)
	self.AddType(&TScopeForm)
	self.AddType(&TSlice)
	self.AddType(&TString)

	self.AddVal("NIL", &TNil, nil)
	self.AddVal("T", &TBool, true)
	self.AddVal("F", &TBool, false)

	self.AddMacro("?:", 2, branchImp)
 	self.AddMacro("call", 0, callImp)
 	self.AddMacro("call:", 1, callArgsImp)
	self.AddMacro("check:", 1, checkImp)
	self.AddMacro("define:", 2, defineImp)
 	self.AddMacro("do:", 1, doImp)
	self.AddMacro("_", 0, dropImp)
	self.AddMacro("..", 0, dupImp)
	self.AddMacro("include:", 1, includeImp)
	self.AddMacro("is", 0, isImp)
	self.AddMacro("\\:", 2, lambdaImp)
	self.AddMacro("let:", 2, letImp)
	self.AddMacro("macro:", 3, macroImp)
	self.AddMacro("method:", 3, methodImp)
	self.AddMacro(",", 0, pairImp)
	self.AddMacro("pause:", 1, pauseImp)
	self.AddMacro("scope:", 1, scopeImp)
	self.AddMacro("thread:", 2, threadImp)
	self.AddMacro("type:", 2, typeDefImp)
	self.AddMacro("use:", 2, useImp)

	self.AddFunction("set")
	self.AddFunction("union")
	
	self.AddMethod("bool", []Arg{AType("val", &TAny)}, []Ret{RType(&TBool)}, boolImp)
	self.AddMethod("clone", []Arg{AType("val", &TAny)}, []Ret{RIndex(0)}, cloneImp)
	self.AddMethod("dump", []Arg{AType("val", &TOptional)}, nil, dumpImp)
	self.AddMethod("=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, eqImp)
	self.AddMethod(">", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, gtImp)
	self.AddMethod(">=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, gteImp)
	self.AddMethod("+", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intAddImp)
	self.AddMethod("*", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intMulImp)
	self.AddMethod("...", []Arg{AType("val", &TInt)}, nil, intSpreadImp)
	self.AddMethod("-", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intSubImp)

	self.AddMethod("isa",
		[]Arg{AType("child", &TMeta), AType("parent", &TMeta)},
		[]Ret{RType(Optional(&TMeta))},
		isaImp)
	
	self.AddMethod("load", []Arg{AType("path", &TString)}, nil, loadImp)
	self.AddMethod("<", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, ltImp)
	self.AddMethod("<=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, lteImp)
	self.AddMethod("...", []Arg{AType("val", &TPair)}, nil, pairSpreadImp)
	self.AddMethod("say", []Arg{AType("val", &TAny)}, nil, sayImp)
	self.AddMethod("length", []Arg{AType("val", &TSlice)}, []Ret{RType(&TInt)}, sliceLengthImp)
	self.AddMethod("peek", []Arg{AType("val", &TSlice)}, []Ret{RType(&TOptional)}, slicePeekImp)
	self.AddMethod("pop", []Arg{AType("val", &TSlice)}, []Ret{RType(&TOptional)}, slicePopImp)
	self.AddMethod("push", []Arg{AType("target", &TSlice), AType("val", &TAny)}, nil, slicePushImp)
	self.AddMethod("...", []Arg{AType("val", &TSlice)}, nil, sliceSpreadImp)
	self.AddMethod("length", []Arg{AType("val", &TString)}, []Ret{RType(&TInt)}, stringLengthImp)
	self.AddMethod("...", []Arg{AType("val", &TString)}, nil, stringSpreadImp)
	self.AddMethod("type", []Arg{AType("val", &TAny)}, []Ret{RType(&TMeta)}, typeImp)
	return self
}
