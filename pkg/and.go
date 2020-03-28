package gfoo

type And struct {
	OpBase
	right []Op
}

func NewAnd(form Form, right []Op) *And {
	op := new(And)
	op.OpBase.Init(form)
	op.right = right
	return op
}

func (self *And) Eval(thread *Thread, registers, stack *Stack) error {
	left := stack.Peek()

	if left == nil {
		return Error(self.form.Pos(), "Missing left operand")
	}

	if !left.Bool() {
		return nil
	}

	stack.Pop()
	
	if err := EvalOps(self.right, thread, registers, stack); err != nil {
		return err
	}

	right := stack.Peek()
	
	if right == nil {
		return Error(self.form.Pos(), "Missing right operand")
	}

	return nil
}
