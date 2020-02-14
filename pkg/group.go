package gfoo

type Group struct {
	FormBase
	forms []Form
}

func NewGroup(forms []Form, pos Pos) *Group {
	return new(Group).Init(forms, pos)
}

func (self *Group) Init(forms []Form, pos Pos) *Group {
	self.FormBase.Init(pos)
	self.forms = forms
	return self
}

func (self *Group) AddForm(form Form) {
	self.forms = append(self.forms, form)
}

func (self *Group) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return scope.Compile(self.forms, out)
}

func (self *Group) Quote(scope *Scope) (Val, error) {
	out := make([]Val, len(self.forms))
	var err error
	
	for i, f := range self.forms {
		if out[i], err = f.Quote(scope); err != nil {
			return NilVal, err
		}
	}
	
	return NewVal(&TSlice, NewSlice(out)), nil
}
