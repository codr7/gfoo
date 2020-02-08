package gfoo

type SliceForm struct {
	Group
}

func NewSliceForm(pos Pos, forms []Form) *SliceForm {
	f := new(SliceForm)
	f.Group.Init(pos, forms)
	return f
}

func (self *SliceForm) Compile(gfoo *GFoo, scope *Scope, in *Forms, out []Op) ([]Op, error) {
	ops, err := gfoo.Compile(self.forms, scope, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewPushSlice(self, ops)), nil
}
