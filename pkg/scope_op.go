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

func (self *ScopeOp) Eval(thread *Thread, registers, stack *Slice) error {
	for _, m := range self.scope.methods {
		for f, i := range m.indexes {
			if i == -1 {
				f.AddMethod(m)
			}
		}
	}

	err := EvalOps(self.body, thread, registers, stack)

	for _, m := range self.scope.methods {
		for f, i := range m.indexes {
			if i > -1 {
				f.RemoveMethod(m)
			}
		}
	}

	return err
}
