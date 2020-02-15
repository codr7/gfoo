package gfoo

type VM struct {
	rootScope Scope
}

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
	return append(out, NewCall(form, nil)), nil
}
	
func letImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	key, ok := in.Pop().(*Id)

	if !ok {
		scope.Error(key.Pos(), "Expected id: %v", key)
	}

	if found := scope.Get(key.name); found == nil {
		scope.Set(key.name, NilVal)
	} else if found.scope != scope {
		found.Init(scope, NilVal)
	} else {
	        return out, scope.Error(key.Pos(), "Duplicate binding: %v", key.name) 
	}
	
	val := in.Pop()
	var err error
	
	if out, err = val.Compile(&NilForms, out, scope); err != nil {
		return out, err
	}
	
	return append(out, NewLet(form, key.name)), nil
}

func threadImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	arg := in.Pop()
	var argOps []Op
	var err error
	
	if argOps, err = arg.Compile(&NilForms, nil, scope); err != nil {
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
	
func NewVM() *VM {
	vm := new(VM)
	vm.rootScope.Init(nil)
	
	vm.AddConst("T", &TBool, true)
	vm.AddConst("F", &TBool, false)

	vm.AddMacro("_", 0, dropImp)
	vm.AddMacro("..", 0, dupImp)
	vm.AddMacro("|", 0, resetImp)
	vm.AddMacro("call", 0, callImp)
	vm.AddMacro("thread:", 1, threadImp)
	vm.AddMacro("let:", 0, letImp)
	vm.AddMacro("type", 0, typeImp)
	return vm
}

func (self *VM) AddConst(name string, dataType Type, data interface{}) {
	self.rootScope.Set(name, NewVal(dataType, data))
}

func (self *VM) AddMacro(name string, argCount int, imp MacroImp) {
	self.AddConst(name, &TMacro, NewMacro(name, argCount, imp))
}

func (self *VM) RootScope() *Scope {
	return &self.rootScope
}
