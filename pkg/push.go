package gfoo

type Push struct {
	OpBase
	val Val
}

func NewPush(form Form, val Val) *Push {
	op := new(Push)
	op.OpBase.Init(form)
	op.val = val
	return op
}

func (self *Push) Eval(thread *Thread, registers []Val, stack *Stack) error {
	stack.Push(self.val)
	return nil
}
