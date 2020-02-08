package gfoo

import (
)

type MacroImp = func(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error)

type Macro struct {
	name string
	argCount int
	imp MacroImp
}

func NewMacro(name string, argCount int, imp MacroImp) *Macro {
	return &Macro{name: name, argCount: argCount, imp: imp}
}

func (self *Macro) Expand(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	if args.Len() < self.argCount {
		gfoo.Error(form.Pos(), "Not enough arguments: %v", self.name)
	}
	
	return self.imp(gfoo, scope, form, args, out)
}

func (self *GFoo) AddMacro(name string, argCount int, imp MacroImp) {
	self.rootScope.Set(name, &TMacro, NewMacro(name, argCount, imp))
}
