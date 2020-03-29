package gfoo

type Dup struct {
	OpBase
}

func NewDup(form Form) *Dup {
	op := new(Dup)
	op.OpBase.Init(form)
	return op
}

func (self *Dup) Eval(thread *Thread, registers []Val, stack *Stack) error {
	v := stack.Peek()
	
	if v == nil {
		return Error(self.form.Pos(), "Nothing to dup")
	}

	stack.Push(*v)
	return nil
}
