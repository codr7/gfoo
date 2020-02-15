package gfoo

type SliceOp struct {
	OpBase
	ops []Op
}

func NewSliceOp(form Form, ops []Op) *SliceOp {
	o := new(SliceOp)
	o.OpBase.Init(form)
	o.ops = ops
	return o
}

func (self *SliceOp) Evaluate(scope *Scope, stack *Slice) error {
	i := stack.Len()

	if err := scope.Evaluate(self.ops, stack); err != nil {
		return err
	}

	stack.Push(NewVal(&TSlice, NewSlice(stack.Cut(i))))
	return nil
}

