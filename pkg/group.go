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

func (self *Group) Compile(vm *VM, scope *Scope, in *Forms, out []Op) ([]Op, error) {
	return vm.Compile(self.forms, scope, out)
}

func (self *Group) Quote() Val {
	out := make([]Val, len(self.forms))

	for i, f := range self.forms {
		out[i] = f.Quote()
	}
	
	return NewVal(&TSlice, NewSlice(out))
}
