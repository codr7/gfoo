package gfoo

type Group struct {
	FormBase
	forms []Form
}

func NewGroup(pos Pos, forms []Form) *Group {
	return new(Group).Init(pos, forms)
}

func (self *Group) Init(pos Pos, forms []Form) *Group {
	self.FormBase.Init(pos)
	self.forms = forms
	return self
}

func (self *Group) AddForm(form Form) {
	self.forms = append(self.forms, form)
}

func (self *Group) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return scope.vm.Compile(self.forms, scope, out)
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
