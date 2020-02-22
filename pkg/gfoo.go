package gfoo

import (
	"math/big"
)

const (
	VersionMajor = 0
	VersionMinor = 6
)

func Init() {
	TAny.Init("Any")
	TBool.Init("Bool", &TAny)
	TFunction.Init("Function", &TAny)
	TId.Init("Id", &TAny)
	TInt.Init("Int", &TAny)
	TLambda.Init("Lambda", &TAny)
	TMacro.Init("Macro", &TAny)
	TMeta.Init("Type", &TAny)
	TMethod.Init("Method", &TAny)
	TPair.Init("Pair", &TAny)
	TScope.Init("Scope", &TAny)
	TString.Init("String", &TAny)
}

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

func dropImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDrop(form)), nil
}

func dupImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDup(form)), nil
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
	
	if out, err = val.Compile(nil, out, scope); err != nil {
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
		a := args.body[i]
		var id *Id
		
		if id, ok = args.body[i].(*Id); !ok {
			return out, scope.Error(a.Pos(), "Invalid argumnet: %v", f)
		}
		
		bodyOps = append(bodyOps, NewLet(id, id.name))
	}

	var err error
	macroScope := scope.Clone()
	
	if bodyOps, err = body.Compile(nil, bodyOps, macroScope); err != nil {
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
		
		if err := scope.Evaluate(bodyOps, &stack); err != nil {
			return out, err
		}

		for _, v := range stack.items {
			in.Push(v.Unquote(scope, form.Pos()))
		}

		return out, nil
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

func resetImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewReset(form)), nil
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

func eqImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	stack.Push(NewVal(&TBool, x.Compare(y) == Eq))
	return nil
}

func gtImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	stack.Push(NewVal(&TBool, x.Compare(y) == Gt))
	return nil
}

func gteImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	stack.Push(NewVal(&TBool, x.Compare(y) >= Eq))
	return nil
}

func intAddImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	var z big.Int
	z.Add(x.data.(*big.Int), y.data.(*big.Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func intMulImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	var z big.Int
	z.Mul(x.data.(*big.Int), y.data.(*big.Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func intSubImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	var z big.Int
	z.Sub(x.data.(*big.Int), y.data.(*big.Int))
	stack.Push(NewVal(&TInt, &z))
	return nil
}

func isImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	stack.Push(NewVal(&TBool, x.Is(y)))
	return nil
}

func loadImp(stack *Slice, scope *Scope, pos Pos) (error) {
	path, _ := stack.Pop()
	return scope.Load(path.data.(string), stack)
}

func ltImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	stack.Push(NewVal(&TBool, x.Compare(y) == Lt))
	return nil
}

func lteImp(stack *Slice, scope *Scope, pos Pos) (error) {
	y, _ := stack.Pop()
	x, _ := stack.Pop()
	stack.Push(NewVal(&TBool, x.Compare(y) <= Eq))
	return nil
}

func typeImp(stack *Slice, scope *Scope, pos Pos) (error) {
	val, _ := stack.Pop();
	stack.Push(NewVal(&TMeta, val.dataType))
	return nil
}

func New() *Scope {
	return new(Scope).InitRoot()
}

func (self *Scope) InitRoot() *Scope {
	self.Init()
	self.AddType(&TAny)
	self.AddType(&TBool)
	self.AddType(&TFunction)
	self.AddType(&TId)
	self.AddType(&TInt)
	self.AddType(&TLambda)
	self.AddType(&TMacro)
	self.AddType(&TMeta)
	self.AddType(&TMethod)
	self.AddType(&TPair)
	self.AddType(&TScope)
	self.AddType(&TString)

	self.AddConst("T", &TBool, true)
	self.AddConst("F", &TBool, false)

	self.AddMacro("?:", 2, branchImp)
 	self.AddMacro("call", 0, callImp)
 	self.AddMacro("call:", 1, callArgsImp)
	self.AddMacro("check:", 1, checkImp)
	self.AddMacro("_", 0, dropImp)
	self.AddMacro("..", 0, dupImp)
	self.AddMacro("\\:", 2, lambdaImp)
	self.AddMacro("let:", 2, letImp)
	self.AddMacro("macro:", 2, macroImp)
	self.AddMacro(",", 0, pairImp)
	self.AddMacro("pause:", 1, pauseImp)
	self.AddMacro("|", 0, resetImp)
	self.AddMacro("thread:", 1, threadImp)

	
	self.AddMethod("=",
		[]Argument{AType("x", &TAny), AType("y", &TAny)},
		[]Result{RType(&TBool)},
		eqImp)

	self.AddMethod(">",
		[]Argument{AType("x", &TAny), AType("y", &TAny)},
		[]Result{RType(&TBool)},
		gtImp)

	self.AddMethod(">=",
		[]Argument{AType("x", &TAny), AType("y", &TAny)},
		[]Result{RType(&TBool)},
		gteImp)

	self.AddMethod("+",
		[]Argument{AType("x", &TInt), AType("y", &TInt)},
		[]Result{RType(&TInt)},
		intAddImp)

	self.AddMethod("*",
		[]Argument{AType("x", &TInt), AType("y", &TInt)},
		[]Result{RType(&TInt)},
		intMulImp)

	self.AddMethod("-",
		[]Argument{AType("x", &TInt), AType("y", &TInt)},
		[]Result{RType(&TInt)},
		intSubImp)

	self.AddMethod("is",
		[]Argument{AType("x", &TAny), AType("y", &TAny)},
		[]Result{RType(&TBool)},
		isImp)

	self.AddMethod("load", []Argument{AType("path", &TString)}, nil, loadImp)

	self.AddMethod("<",
		[]Argument{AType("x", &TAny), AType("y", &TAny)},
		[]Result{RType(&TBool)},
		ltImp)
	
	self.AddMethod("<=",
		[]Argument{AType("x", &TAny), AType("y", &TAny)},
		[]Result{RType(&TBool)},
		lteImp)

	self.AddMethod("type", []Argument{AType("val", &TAny)}, []Result{RType(&TMeta)}, typeImp)
	return self
}

