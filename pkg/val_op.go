package gfoo

type ValOp struct {
	OpBase
}

func NewValOp(form Form) *ValOp {
	op := new(ValOp)
	op.OpBase.Init(form)
	return op
}

func (self *ValOp) Eval(scope *Scope, stack *Slice) error {
	stack.Push(scope.val)
	return nil
}

