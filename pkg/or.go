package gfoo

type Or struct {
	OpBase
	right []Op
}

func NewOr(form Form, right []Op) *Or {
	op := new(Or)
	op.OpBase.Init(form)
	op.right = right
	return op
}

func (self *Or) Eval(scope *Scope, stack *Slice) error {
	left := stack.Peek()

	if left == nil {
		return scope.Error(self.form.Pos(), "Missing left operand")
	}

	if left.Bool() {
		return nil
	}

	stack.Pop()
	
	if err := scope.EvalOps(self.right, stack); err != nil {
		return err
	}

	right := stack.Peek()
	
	if right == nil {
		return scope.Error(self.form.Pos(), "Missing right operand")
	}

	return nil
}
