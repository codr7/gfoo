package gfoo

type Group struct {
	FormBase
	items []Form
}

func NewGroup(pos Pos, items []Form) *Group {
	f := new(Group)
	f.FormBase.Init(pos)
	f.items = items
	return f
}

func (self *Group) Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error) {
	ops, err := gfoo.Compile(self.items, scope, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewPushSlice(self, ops)), nil
}

func (self *Group) Quote() Val {
	v := make([]Val, len(self.items))

	for i, f := range self.items {
		v[i] = f.Quote()
	}
	
	return NewVal(&Slice, v)
}
