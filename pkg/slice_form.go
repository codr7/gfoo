package gfoo

type SliceForm struct {
	FormBase
	body []Form
}

func NewSliceForm(body []Form, pos Pos) *SliceForm {
	return new(SliceForm).Init(body, pos)
}

func (self *SliceForm) Init(body []Form, pos Pos) *SliceForm {
	self.FormBase.Init(pos)
	self.body = body
	return self
}

func (self *SliceForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	ops, err := scope.Compile(self.body, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewSliceOp(self, ops)), nil
}

func (self *SliceForm) Do(action func(Form) error) error {
	for _, f := range self.body {
		if err := f.Do(action); err != nil {
			return err
		}
	}

	return nil
}

func (self *SliceForm) Quote(scope *Scope) (Val, error) {
	ops, err := scope.Compile(self.body, nil)

	if err != nil {
		return NilVal, err
	}

	v := NewSlice(nil)
	
	if err = scope.Evaluate(ops, v); err != nil {
		return NilVal, err
	}
		
	return NewVal(&TSlice, v), nil
}
