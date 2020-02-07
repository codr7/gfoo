package gfoo

type Group struct {
	forms []Form
}

func NewGroup(forms []Form) *Group {
	return &Group{forms: forms}
}

func (grp *Group) Quote() Value {
	v := make([]Value, len(grp.forms))

	for i, f := range grp.forms {
		v[i] = f.Quote()
	}
	
	return NewValue(&Slice, v)
}
