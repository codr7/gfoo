package gfoo

type Is struct {
	OpBase
}

func NewIs(form Form) *Is {
	op := new(Is)
	op.OpBase.Init(form)
	return op
}

func (self *Is) Evaluate(scope *Scope, stack *Slice) error {
	var left, right Val
	var ok bool
	
	if right, ok = stack.Pop(); !ok {
		return scope.Error(self.form.Pos(), "Missing right value")
	}

	if left, ok = stack.Pop(); !ok {
		return scope.Error(self.form.Pos(), "Missing left value")
	}

	stack.Push(NewVal(&TBool, left.data == right.data))
	return nil
}
