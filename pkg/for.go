package gfoo

type For struct {
	OpBase
	body []Op
}

func NewFor(form Form, body []Op) *For {
	op := new(For)
	op.OpBase.Init(form)
	op.body = body
	return op
}

func (self *For) Eval(scope *Scope, stack *Slice) error {
	v := stack.Pop()

	if v == nil {
		return scope.Error(self.form.Pos(), "Missing value")
	}

	in, err := v.Iter(scope, self.form.Pos())

	if err != nil {
		return err
	}
	
	return in.For(func(v Val, scope *Scope, pos Pos) error {
		scope.val.Push(v)
		defer scope.val.Pop()
		return scope.EvalOps(self.body, stack)
	}, scope, self.form.Pos())
}
