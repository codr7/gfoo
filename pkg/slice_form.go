package gfoo

type SliceForm struct {
	FormBase
	forms []Form
}

func NewSliceForm(forms []Form, pos Pos) *SliceForm {
	return new(SliceForm).Init(forms, pos)
}

func (self *SliceForm) Init(forms []Form, pos Pos) *SliceForm {
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
