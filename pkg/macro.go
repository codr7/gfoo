package gfoo

import (
)

type MacroImp = func(vm *VM, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error)

type Macro struct {
	name string
	argCount int
	imp MacroImp
}

func NewMacro(name string, argCount int, imp MacroImp) *Macro {
	return &Macro{name: name, argCount: argCount, imp: imp}
}

func (self *Macro) Expand(vm *VM, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	if args.Len() < self.argCount {
		vm.Error(form.Pos(), "Not enough arguments: %v", self.name)
	}
	
	return self.imp(vm, scope, form, args, out)
}
