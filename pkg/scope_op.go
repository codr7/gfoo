package gfoo

import (
	//"fmt"
)

type ScopeOp struct {
	OpBase
	body []Op
}

func NewScopeOp(form Form, body []Op) *ScopeOp {
	o := new(ScopeOp)
	o.OpBase.Init(form)
	o.body = body
	return o
}

func (self *ScopeOp) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	return vm.Evaluate(self.body, stack, scope.Clone(true))
}

