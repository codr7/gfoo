package gfoo

import (
	"math/big"
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

	f = in.Pop()
	var body *ScopeForm

	if body, ok = f.(*ScopeForm); !ok {
		return out, scope.Error(form.Pos(), "Invalid body: %v", f)
	}

	var args []Arg
	
	for i := 0; i < len(argsForm.body)-1; i++ {
		anf := argsForm.body[i]
		an := anf.(*Id).name
		i++

		atnf := argsForm.body[i]
		atn := atnf.(*Id).name
		atb := scope.Get(atn)

		if atb == nil {
			return out, scope.Error(atnf.Pos(), "Type not found: %v", atn)
		}

		if atb.val.dataType != &TMeta {
			return out, scope.Error(atnf.Pos(), "Expected type: %v", atb.val.dataType.Name())
		}

		args = append(args, AType(an, atb.val.data.(Type)))
	}

	var bodyOps []Op

	for i := len(args)-1; i >= 0; i-- {
		bodyOps = append(bodyOps, NewLet(argsForm, args[i].name))
	}

	var rets []Ret

	for _, f := range retsForm.body {
		rtn := f.(*Id).name
		rtb := scope.Get(rtn)

		if rtb == nil {
			return out, scope.Error(f.Pos(), "Type not found: %v", rtn)
		}

		if rtb.val.dataType != &TMeta {
			return out, scope.Error(f.Pos(), "Expected type: %v", rtb.val.dataType.Name())
		}

		rets = append(rets, RType(rtb.val.data.(Type)))
	}
	
	methodScope := scope.Clone()
	var err error
	
	if bodyOps, err = methodScope.Compile(body.body, bodyOps); err != nil {
		return out, err
	}

	scope.AddMethod(id.name, args, rets, func(stack *Slice, scope *Scope, pos Pos) error {
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
		func(stack *Slice, scope *Scope, pos Pos) error {
			stack.Pop()
			v := stack.Peek()
			v.dataType = t
			return nil
		})
	
	return out, nil
}

func boolImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TBool, stack.Pop().Bool()))
	return nil
}

func cloneImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(stack.Pop().Clone())
	return nil
}

func dumpImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Pop().Dump(os.Stdout)
	os.Stdout.WriteString("\n")
	return nil
}

func eqImp(stack *Slice, scope *Scope, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Eq))
	return nil
}

func gtImp(stack *Slice, scope *Scope, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Gt))
	return nil
}

func gteImp(stack *Slice, scope *Scope, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) >= Eq))
	return nil
}

func intAddImp(stack *Slice, scope *Scope, pos Pos) error {
	var z big.Int
	z.Add(stack.Pop().data.(*big.Int), stack.Pop().data.(*big.Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func intMulImp(stack *Slice, scope *Scope, pos Pos) error {
	var z big.Int
	z.Mul(stack.Pop().data.(*big.Int), stack.Pop().data.(*big.Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func intSubImp(stack *Slice, scope *Scope, pos Pos) error {
	y := stack.Pop()
	var z big.Int
	z.Sub(stack.Pop().data.(*big.Int), y.data.(*big.Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func loadImp(stack *Slice, scope *Scope, pos Pos) error {
	return scope.Load(stack.Pop().data.(string), stack)
}

func ltImp(stack *Slice, scope *Scope, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) == Lt))
	return nil
}

func lteImp(stack *Slice, scope *Scope, pos Pos) error {
	y := stack.Pop()
	stack.Push(NewVal(&TBool, stack.Pop().Compare(*y) <= Eq))
	return nil
}

func sayImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Pop().Print(os.Stdout)
	os.Stdout.WriteString("\n")
	return nil
}

func sliceLengthImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TInt, big.NewInt(int64(stack.Pop().data.(*Slice).Len()))))
	return nil
}

func stringLengthImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TInt, big.NewInt(int64(len(stack.Pop().data.(string))))))
	return nil
}

func typeImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TMeta, stack.Pop().dataType))
	return nil
}

func (self *Scope) InitAbc() *Scope {
	self.AddType(&TAny)
	self.AddType(&TBool)
	self.AddType(&TFunction)
	self.AddType(&TId)
	self.AddType(&TInt)
	self.AddType(&TLambda)
	self.AddType(&TMacro)
	self.AddType(&TMeta)
	self.AddType(&TMethod)
	self.AddType(&TNil)
	self.AddType(&TNumber)
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

	self.AddFunction("set")
	self.AddFunction("union")
	
	self.AddMethod("bool", []Arg{AType("val", &TAny)}, []Ret{RType(&TBool)}, boolImp)
	self.AddMethod("clone", []Arg{AType("val", &TAny)}, []Ret{RIndex(0)}, cloneImp)
	self.AddMethod("dump", []Arg{AType("val", &TAny)}, nil, dumpImp)
	self.AddMethod("=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, eqImp)
	self.AddMethod(">", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, gtImp)
	self.AddMethod(">=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, gteImp)
	self.AddMethod("+", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intAddImp)
	self.AddMethod("*", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intMulImp)
	self.AddMethod("-", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, intSubImp)
	self.AddMethod("load", []Arg{AType("path", &TString)}, nil, loadImp)
	self.AddMethod("<", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, ltImp)
	self.AddMethod("<=", []Arg{AType("x", &TAny), AType("y", &TAny)}, []Ret{RType(&TBool)}, lteImp)
	self.AddMethod("say", []Arg{AType("val", &TAny)}, nil, sayImp)
	self.AddMethod("length", []Arg{AType("val", &TSlice)}, []Ret{RType(&TInt)}, sliceLengthImp)
	self.AddMethod("length", []Arg{AType("val", &TString)}, []Ret{RType(&TInt)}, stringLengthImp)
	self.AddMethod("type", []Arg{AType("val", &TAny)}, []Ret{RType(&TMeta)}, typeImp)
	return self
}
