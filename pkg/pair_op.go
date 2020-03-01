package gfoo

type PairOp struct {
	OpBase
}

func NewPairOp(form Form) *PairOp {
	op := new(PairOp)
	op.OpBase.Init(form)
	return op
}

func (self *PairOp) Eval(scope *Scope, stack *Slice) error {
	var left, right *Val
		
	if right = stack.Pop(); right == nil {
		return scope.Error(self.form.Pos(), "Missing right")
	}
	
	if left = stack.Pop(); left == nil {
		return scope.Error(self.form.Pos(), "Missing left")
	}

	stack.Push(NewVal(&TPair, NewPair(*left, *right)))
	return nil
}
