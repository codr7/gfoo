package gfoo

type SliceForm struct {
	Group
}

func NewSliceForm(pos Pos, forms []Form) *SliceForm {
	f := new(SliceForm)
	f.Group.Init(pos, forms)
	return f
}

func (self *SliceForm) Compile(vm *VM, scope *Scope, in *Forms, out []Op) ([]Op, error) {
	ops, err := vm.Compile(self.forms, scope, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewPushSlice(self, ops)), nil
}
