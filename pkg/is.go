package gfoo

type Is struct {
	OpBase
	x *Val
}

func NewIs(form Form, x *Val) *Is {
	op := new(Is)
	op.OpBase.Init(form)
	op.x = x
	return op
}

func (self *Is) Eval(scope *Scope, stack *Slice) error {
	y := stack.Pop()

	if y == nil {
		return scope.Error(self.form.Pos(), "Missing right value")
	}

	x := self.x

	if x == nil {			
		if x = stack.Pop(); x == nil {
			return scope.Error(self.form.Pos(), "Missing left value")
		}
	}

	stack.Push(NewVal(&TBool, x.Is(*y)))
	return nil
}

