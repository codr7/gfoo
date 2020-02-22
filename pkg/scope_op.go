package gfoo

import (
	//"fmt"
)

type ScopeOp struct {
	OpBase
	body []Op
	scope *Scope
}

func NewScopeOp(form Form, body []Op, scope *Scope) *ScopeOp {
	op := new(ScopeOp)
	op.OpBase.Init(form)
	op.body = body
	op.scope = scope
	return op
}

func (self *ScopeOp) Evaluate(scope *Scope, stack *Slice) error {
	return self.scope.Clone().Evaluate(self.body, stack)
}

