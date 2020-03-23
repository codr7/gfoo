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

func (self *For) Eval(thread *Thread, registers, stack *Slice) error {
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing value")
	}

	in, err := v.Iter(self.form.Pos())

	if err != nil {
		return err
	}
	
	return in.For(func(v Val, thread *Thread, pos Pos) error {
		stack.Push(v)
		return EvalOps(self.body, thread, registers, stack)
	}, thread, self.form.Pos())
}
