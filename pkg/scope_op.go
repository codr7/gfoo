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

func (self *ScopeOp) Eval(scope *Scope, stack *Slice) error {
	s := self.scope

	if s == nil {
		sv := stack.Pop()
		
		if sv == nil {
			return scope.Error(self.form.Pos(), "Missing scope")
		}

		if sv.dataType != &TScope {
			return scope.Error(self.form.Pos(), "Expected scope: %v", sv)
		}

		s = sv.data.(*Scope)
	}
	
	for _, m := range s.methods {
		if m.index == -1 {
			m.function.AddMethod(m)
		}
	}

	es := s

	if self.scope != nil {
		es = es.Clone().Extend(scope)
	}
	
	err := es.EvalOps(self.body, stack)

	for _, m := range s.methods {
		m.function.RemoveMethod(m)
	}

	return err
}
