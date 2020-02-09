package gfoo

type GetType struct {
	OpBase
}

func NewGetType(form Form) *GetType {
	o := new(GetType)
	o.OpBase.Init(form)
	return o
}

func (self *GetType) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	v := stack.Peek()

	if v == nil {
		return vm.Error(self.form.Pos(), "Missing value")
	}

	v.data, v.dataType = v.dataType, &TMeta
	return nil
}
