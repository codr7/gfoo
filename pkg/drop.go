package gfoo

type Drop struct {
	OpBase
}

func NewDrop(form Form) *Drop {
	o := new(Drop)
	o.OpBase.Init(form)
	return o
}

func (self *Drop) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	if v := stack.Pop(); v == nil {
		return vm.Error(self.form.Pos(), "Nothing to drop")
	}
	
	return nil
}
