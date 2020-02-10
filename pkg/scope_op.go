package gfoo

import (
	//"fmt"
)

type ScopeOp struct {
	OpBase
	ops []Op
}

func NewScopeOp(form Form, ops []Op) *ScopeOp {
	o := new(ScopeOp)
	o.OpBase.Init(form)
	o.ops = ops
	return o
}

func (self *ScopeOp) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	return vm.Evaluate(self.ops, stack, scope.Clone())
}

