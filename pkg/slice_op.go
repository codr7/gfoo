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

func (self *SliceOp) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	i := stack.Len()

	if err := vm.Evaluate(self.ops, stack, scope); err != nil {
		return err
	}

	stack.Push(NewVal(&TSlice, NewSlice(stack.Cut(i))))
	return nil
}

