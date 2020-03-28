package gfoo

type PairOp struct {
	OpBase
	left, right []Op
}

func NewPairOp(form Form, left, right []Op) *PairOp {
	op := new(PairOp)
	op.OpBase.Init(form)
	op.left = left
	op.right = right
	return op
}

func (self *PairOp) Eval(thread *Thread, registers, stack *Stack) error {
	var left, right *Val
	
	if err := EvalOps(self.left, thread, registers, stack); err != nil {
		return err
	}
	
	if left = stack.Pop(); left == nil {
		return Error(self.form.Pos(), "Missing left")
	}
	
	if err := EvalOps(self.right, thread, registers, stack); err != nil {
		return err
	}

	if right = stack.Pop(); right == nil {
		return Error(self.form.Pos(), "Missing right")
	}

	stack.Push(NewVal(&TPair, NewPair(*left, *right)))
	return nil
}
