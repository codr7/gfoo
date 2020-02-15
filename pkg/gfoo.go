package gfoo

import (
	//"fmt"
)

const (
	VersionMajor = 0
	VersionMinor = 3
)

func dropImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDrop(form)), nil
}

func dupImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewDup(form)), nil
}

func resetImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewReset(form)), nil
}

func callImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	return append(out, NewCall(form, nil, nil)), nil
}

func callArgsImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error){
	arg := in.Pop()
	var argOps []Op
	var err error
	
	if argOps, err = arg.Compile(NewForms(nil), nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewCall(form, nil, argOps)), nil
}

func lambdaImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var args *Group
	var ok bool
	
	if args, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid argument list: %v", f)
	}

	f = in.Pop()
	var bodyForm *ScopeForm

	if bodyForm, ok = f.(*ScopeForm); !ok {
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
	
	if bodyOps, err = scope.Compile(bodyForm.body, bodyOps); err != nil {
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
		scope.Set(id.name, NilVal)
	} else if found.scope != scope {
		found.Init(scope, NilVal)
	} else {
	        return out, scope.Error(id.Pos(), "Duplicate binding: %v", id.name) 
	}
	
	val := in.Pop()
	var err error
	
	if out, err = val.Compile(NewForms(nil), out, scope); err != nil {
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
	var bodyForm *ScopeForm

	if bodyForm, ok = f.(*ScopeForm); !ok {
		return out, scope.Error(form.Pos(), "Invalid body: %v", f)
	}

	body := bodyForm.body
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
	
	if bodyOps, err = macroScope.Compile(body, bodyOps); err != nil {
		return out, err
	}
	
	scope.AddMacro(id.name, len(args.body), func(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
		var stack Slice

		for i := 0; i < len(args.body); i++ {
			var v Val
			
			if v, err = in.Pop().Quote(scope); err != nil {
				return out, err
			}
			
			stack.Push(v)
		}

		scope = macroScope.Clone()
		
		if err := scope.Evaluate(bodyOps, &stack); err != nil {
			return out, err
		}

		for _, v := range stack.items {
			in.Push(v.Unquote(macroScope, form.Pos()))
		}

		return out, nil
	})
		
	return out, nil
}

func pauseImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	result := in.Pop()
	var resultOps []Op
	var err error
	
	if resultOps, err = result.Compile(NewForms(nil), nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewPause(form, resultOps)), nil
}

func threadImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	arg := in.Pop()
	var argOps []Op
	var err error
	
	if argOps, err = arg.Compile(NewForms(nil), nil, scope); err != nil {
		return out, err
	}

	body := in.Pop()
	var bodyForms []Form

	if f, ok := body.(*ScopeForm); ok {
		bodyForms = f.body
	} else {
		bodyForms = append(bodyForms, body)
	}
	
	var bodyOps []Op
	
	if bodyOps, err = scope.Compile(bodyForms, nil); err != nil {
		return out, err
	}
	
	return append(out, NewThreadOp(form, argOps, bodyOps)), nil
}

func typeImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewTypeOp(form)), nil
}
	
func New() *Scope {
	return new(Scope).InitRoot()
}

func (self *Scope) InitRoot() *Scope {
	self.Init(nil)
	self.AddConst("T", &TBool, true)
	self.AddConst("F", &TBool, false)

	self.AddMacro("_", 0, dropImp)
	self.AddMacro("..", 0, dupImp)
	self.AddMacro("|", 0, resetImp)
 	self.AddMacro("call", 0, callImp)
 	self.AddMacro("call:", 0, callArgsImp)
	self.AddMacro("lambda:", 2, lambdaImp)
	self.AddMacro("let:", 2, letImp)
	self.AddMacro("macro:", 2, macroImp)
	self.AddMacro("pause:", 1, pauseImp)
	self.AddMacro("thread:", 1, threadImp)
	self.AddMacro("type", 0, typeImp)
	return self
}

