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

func (self *SliceForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	ops, err := scope.Compile(self.forms, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewSliceOp(self, ops)), nil
}

func (self *SliceForm) Quote(scope *Scope) (Val, error) {
	ops, err := scope.Compile(self.forms, nil)

	if err != nil {
		return NilVal, err
	}

	v := NewSlice(nil)
	
	if err = scope.Evaluate(ops, v); err != nil {
		return NilVal, err
	}
		
	return NewVal(&TSlice, v), nil
}
