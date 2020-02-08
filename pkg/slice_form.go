package gfoo

type SliceForm struct {
	FormBase
	items []Form
}

func NewSliceForm(pos Pos, items []Form) *SliceForm {
	f := new(SliceForm)
	f.FormBase.Init(pos)
	f.items = items
	return f
}

func (self *SliceForm) Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error) {
	ops, err := gfoo.Compile(self.items, scope, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewPushSlice(self, ops)), nil
}

func (self *SliceForm) Quote() Val {
	v := make([]Val, len(self.items))

	for i, f := range self.items {
		v[i] = f.Quote()
	}
	
	return NewVal(&Slice, v)
}
