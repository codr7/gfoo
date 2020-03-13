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

func (self *PairOp) Eval(scope *Scope, stack *Slice) error {
	var left, right *Val
	
	if err := scope.EvalOps(self.left, stack); err != nil {
		return err
	}
	
	if left = stack.Pop(); left == nil {
		return scope.Error(self.form.Pos(), "Missing left")
	}
	
	if err := scope.EvalOps(self.right, stack); err != nil {
		return err
	}

	if right = stack.Pop(); right == nil {
		return scope.Error(self.form.Pos(), "Missing right")
	}

	stack.Push(NewVal(&TPair, NewPair(*left, *right)))
	return nil
}
