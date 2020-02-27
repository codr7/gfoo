package gfoo

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
	for _, m := range self.scope.methods {
		if m.index == -1 {
			m.function.AddMethod(m)
		}
	}

	
	err := self.scope.Clone().Extend(scope).Evaluate(self.body, stack)

	for _, m := range self.scope.methods {
		m.function.RemoveMethod(m)
	}

	return err
}
