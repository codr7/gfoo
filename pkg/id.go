package gfoo

type Id struct {
	FormBase
	name string
}

func NewId(name string, pos Pos) *Id {
	f := new(Id)
	f.FormBase.Init(pos)
	f.name = name
	return f
}

func (self *Id) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	if b := scope.Get(self.name); b != nil && b.val != NilVal {
		v := &b.val
		
		if v.dataType == &TMacro {
			return v.data.(*Macro).Expand(self, in, out, scope)
		}
		
		return append(out, NewPush(self, *v)), nil
	}

	return append(out, NewGet(self, self.name)), nil
}

func (self *Id) Do(action func(Form) error) error {
	return action(self)
}

func (self *Id) Quote(scope *Scope) (Val, error) {
	return NewVal(&TId, self.name), nil
}
