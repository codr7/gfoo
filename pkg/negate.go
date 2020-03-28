package gfoo

type Negate struct {
	OpBase
}

func NewNegate(form Form) *Negate {
	op := new(Negate)
	op.OpBase.Init(form)
	return op
}

func (self *Negate) Eval(thread *Thread, registers, stack *Stack) error {
	v := stack.Peek()
	
	if v == nil {
		return Error(self.form.Pos(), "Missing value")
	}

	v.Negate()
	return nil
}
