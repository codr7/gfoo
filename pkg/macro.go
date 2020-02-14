package gfoo

import (
)

type MacroImp = func(form Form, args *Forms, out []Op, scope *Scope) ([]Op, error)

type Macro struct {
	name string
	argCount int
	imp MacroImp
}

func NewMacro(name string, argCount int, imp MacroImp) *Macro {
	return &Macro{name: name, argCount: argCount, imp: imp}
}

func (self *Macro) Expand(form Form, args *Forms, out []Op, scope *Scope) ([]Op, error) {
	if args.Len() < self.argCount {
		scope.vm.Error(form.Pos(), "Not enough arguments: %v", self.name)
	}
	
	return self.imp(form, args, out, scope)
}
