package gfoo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type VM struct {
	Debug bool	
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
		scope.vm.Error(key.Pos(), "Expected id: %v", key)
	}

	if found := scope.Get(key.name); found == nil {
		scope.Set(key.name, NilVal)
	} else if found.scope != scope {
		found.Init(scope, NilVal)
	} else {
	        return out, scope.vm.Error(key.Pos(), "Duplicate binding: %v", key.name) 
	}
	
	val := in.Pop()
	
	if id, ok := val.(*Id); !ok || id.name != "_" {
		var err error

		if out, err = val.Compile(&NilForms, out, scope); err != nil {
			return out, err
		}
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
	vm.rootScope.Init(vm, nil)
	
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

func (self *VM) Error(pos Pos, spec string, args...interface{}) error {
	msg := fmt.Sprintf("Error in '%v', line %v, column %v: %v ", 
		pos.source, pos.line, pos.column, fmt.Sprintf(spec, args...))

	if self.Debug {
		panic(msg)
	}

	return errors.New(msg)
}

func (self *VM) Parse(in *bufio.Reader, pos *Pos, out []Form) ([]Form, error) {
	var f Form
	var err error
	
	for {
		if err = skipSpace(in, pos); err == nil {
			f, err = self.parseForm(in, pos)
		}

		if err == io.EOF {
			break
		}

		if err != nil {			
			return out, err
		}

		out = append(out, f)
	}
	
	return out, nil
}

func (self *VM) RootScope() *Scope {
	return &self.rootScope
}
