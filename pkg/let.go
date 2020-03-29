package gfoo

type Let struct {
	OpBase
	key string
	index int
}

func NewLet(form Form, key string, index int) *Let {
	op := new(Let)
	op.OpBase.Init(form)
	op.key = key
	op.index = index
	return op
}

func (self *Let) Eval(thread *Thread, registers []Val, stack *Stack) error {	
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing value: %v", self.key)
	}

	registers[self.index] = *v
	return nil
}

