package gfoo

import (
	//"fmt"
)

type ScopeOp struct {
	OpBase
	body []Op
}

func NewScopeOp(form Form, body []Op) *ScopeOp {
	op := new(ScopeOp)
	op.OpBase.Init(form)
	op.body = body
	return op
}

func (self *ScopeOp) Evaluate(scope *Scope, stack *Slice) error {
	return scope.Clone().Evaluate(self.body, stack)
}

