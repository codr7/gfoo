package gfoo

type For struct {
	OpBase
	body []Op
	scope *Scope
}

func NewFor(form Form, body []Op, scope *Scope) *For {
	op := new(For)
	op.OpBase.Init(form)
	op.body = body
	op.scope = scope
	return op
}

func (self *For) Eval(scope *Scope, stack *Slice) error {
	in := stack.Pop()

	if in == nil {
		return scope.Error(self.form.Pos(), "Missing value")
	}

	if self.scope != nil {
		scope = self.scope.Clone().Extend(scope)
	}
	
	return in.For(func(v Val) error {
		scope.val.Push(v)
		defer scope.val.Pop()
		return scope.EvalOps(self.body, stack)
	}, scope, self.form.Pos())
}

