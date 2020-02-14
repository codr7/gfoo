package gfoo

type SliceForm struct {
	FormBase
	forms []Form
}

func NewSliceForm(pos Pos, forms []Form) *SliceForm {
	return new(SliceForm).Init(pos, forms)
}

func (self *SliceForm) Init(pos Pos, forms []Form) *SliceForm {
	self.FormBase.Init(pos)
	self.forms = forms
	return self
}

func (self *SliceForm) Compile(in *Forms, out []Op, vm *VM, scope *Scope) ([]Op, error) {
	ops, err := vm.Compile(self.forms, scope, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewSliceOp(self, ops)), nil
}

func (self *SliceForm) Quote(vm *VM, scope *Scope) (Val, error) {
	ops, err := vm.Compile(self.forms, scope, nil)

	if err != nil {
		return NilVal, err
	}

	v := NewSlice(nil)
	
	if err = vm.Evaluate(ops, v, scope); err != nil {
		return NilVal, err
	}
		
	return NewVal(&TSlice, v), nil
}
