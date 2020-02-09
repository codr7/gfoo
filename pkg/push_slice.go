package gfoo

type PushSlice struct {
	OpBase
	ops []Op
}

func NewPushSlice(form Form, ops []Op) *PushSlice {
	o := new(PushSlice)
	o.OpBase.Init(form)
	o.ops = ops
	return o
}

func (self *PushSlice) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	i := stack.Len()

	if err := vm.Evaluate(self.ops, stack, scope); err != nil {
		return err
	}

	stack.Push(&TSlice, NewSlice(stack.Cut(i)))
	return nil
}

