package gfoo

type ScopeOp struct {
	OpBase
	body []Op
	scope *Scope
	pop bool
}

func NewScopeOp(form Form, body []Op, scope *Scope, pop bool) *ScopeOp {
	op := new(ScopeOp)
	op.OpBase.Init(form)
	op.body = body
	op.scope = scope
	op.pop = pop
	return op
}

func (self *ScopeOp) Eval(scope *Scope, stack *Slice) error {
	var methodScope *Scope
	
	if self.pop {
		v := stack.Pop()
		
		if v == nil {
			return scope.Error(self.form.Pos(), "Missing scope")
		}

		if v.dataType != &TScope {
			return scope.Error(self.form.Pos(), "Expected scope: %v", v)
		}

		scope = v.data.(*Scope).Extend(self.scope)
		methodScope = scope
	} else {
		scope = self.scope.Clone().Extend(scope)
		methodScope = self.scope
	}
	
	for _, m := range methodScope.methods {
		for f, i := range m.indexes {
			if i == -1 {
				f.AddMethod(m)
			}
		}
	}

	err := scope.EvalOps(self.body, stack)

	for _, m := range methodScope.methods {
		for f, i := range m.indexes {
			if i > -1 {
				f.RemoveMethod(m)
			}
		}
	}

	return err
}
