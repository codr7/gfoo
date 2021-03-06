package gfoo

type Drop struct {
	OpBase
}

func NewDrop(form Form) *Drop {
	op := new(Drop)
	op.OpBase.Init(form)
	return op
}

func (self *Drop) Eval(thread *Thread, registers []Val, stack *Stack) error {
	if stack.Pop() == nil {
		return Error(self.form.Pos(), "Nothing to drop")
	}
	
	return nil
}
