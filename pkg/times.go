package gfoo

type Times struct {
	OpBase
	body []Op
}

func NewTimes(form Form, body []Op) *Times {
	op := new(Times)
	op.OpBase.Init(form)
	op.body = body
	return op
}

func (self *Times) Eval(thread *Thread, registers, stack *Stack) error {
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing repetitions")
	}

	if v.dataType != &TInt {
		return Error(self.form.Pos(), "Expected integer: %v", v.dataType)
	}

	for i := Int(0); i < v.data.(Int); i++ {
		if err := EvalOps(self.body, thread, registers, stack); err != nil {
			return err
		}
	}

	return nil
}
