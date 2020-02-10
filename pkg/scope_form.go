package gfoo

import (
	//"fmt"
)

type ScopeForm struct {
	FormBase
	forms []Form
}

func NewScopeForm(pos Pos, forms []Form) *ScopeForm {
	return new(ScopeForm).Init(pos, forms)
}

func (self *ScopeForm) Init(pos Pos, forms []Form) *ScopeForm {
	self.FormBase.Init(pos)
	self.forms = forms
	return self
}

func (self *ScopeForm) Compile(vm *VM, scope *Scope, in *Forms, out []Op) ([]Op, error) {
	ops, err := vm.Compile(self.forms, scope.Clone(), nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewScopeOp(self, ops)), nil
}

func (self *ScopeForm) Quote(vm *VM, scope *Scope) (Val, error) {
	scope = scope.Clone()
	ops, err := vm.Compile(self.forms, scope, nil)

	if err != nil {
		return NilVal, err
	}

	return NewVal(&TLambda, NewLambda(self.forms, ops, scope)), nil
}
