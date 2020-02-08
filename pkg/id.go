package gfoo

type Id struct {
	FormBase
	name string
}

func NewId(pos Pos, name string) *Id {
	f := new(Id)
	f.FormBase.Init(pos)
	f.name = name
	return f
}

func (self *Id) Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error) {
	if b := scope.Get(self.name); b != nil {
		return append(out, NewPush(self, b.val)), nil
	}

	return append(out, NewGet(self, self.name)), nil
}

func (self *Id) Quote() Val {
	return NewVal(&Symbol, self.name)
}
