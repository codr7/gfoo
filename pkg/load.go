package gfoo

type Load struct {
	OpBase
	key string
	index int
}

func NewLoad(form Form, key string, index int) *Load {
	op := new(Load)
	op.OpBase.Init(form)
	op.key = key
	op.index = index
	return op
}

func (self *Load) Eval(thread *Thread, registers, stack *Slice) error {
	stack.Push(registers.items[self.index])
	return nil
}
