package gfoo

type StackOp struct {
	OpBase
	body []Op
}

func NewStackOp(form Form, body []Op) *StackOp {
	op := new(StackOp)
	op.OpBase.Init(form)
	op.body = body
	return op
}

func (self *StackOp) Eval(thread *Thread, registers, stack *Stack) error {
	s := NewStack(nil)
	
	if err := EvalOps(self.body, thread, registers, s); err != nil {
		return err
	}

	stack.Push(NewVal(&TStack, s))
	return nil
}

