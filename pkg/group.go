package gfoo

type Group struct {
	FormBase
	body []Form
}

func NewGroup(body []Form, pos Pos) *Group {
	return new(Group).Init(body, pos)
}

func (self *Group) Init(body []Form, pos Pos) *Group {
	self.FormBase.Init(pos)
	self.body = body
	return self
}

func (self *Group) Push(form Form) {
	self.body = append(self.body, form)
}

func (self *Group) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return scope.Compile(self.body, out)
}

func (self *Group) Do(action func(Form) error) error {
	for _, f := range self.body {
		if err := f.Do(action); err != nil {
			return err
		}
	}

	return nil
}

func (self *Group) Quote(scope *Scope) (Val, error) {
	out := make([]Val, len(self.body))
	var err error
	
	for i, f := range self.body {
		if out[i], err = f.Quote(scope); err != nil {
			return NilVal, err
		}
	}
	
	return NewVal(&TSlice, NewSlice(out)), nil
}
