package gfoo

type PairOp struct {
	OpBase
}

func NewPairOp(form Form) *PairOp {
	op := new(PairOp)
	op.OpBase.Init(form)
	return op
}

func (self *PairOp) Evaluate(scope *Scope, stack *Slice) error {
	var left, right Val
	var ok bool
		
	if right, ok = stack.Pop(); !ok {
		return scope.Error(self.form.Pos(), "Missing right")
	}
	
	if left, ok = stack.Pop(); !ok {
		return scope.Error(self.form.Pos(), "Missing left")
	}

	stack.Push(NewVal(&TPair, NewPair(left, right)))
	return nil
}
