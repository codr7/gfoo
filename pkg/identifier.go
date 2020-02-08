package gfoo

type Identifier struct {
	FormBase
	name string
}

func NewIdentifier(pos Position, name string) *Identifier {
	f := new(Identifier)
	f.FormBase.Init(pos)
	f.name = name
	return f
}

func (self *Identifier) Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error) {
	if v := scope.Get(self.name); v != nil {
		return append(out, NewPush(self, *v)), nil
	}

	return append(out, NewGet(self, self.name)), nil
}

func (self *Identifier) Quote() Value {
	return NewValue(&Symbol, self.name)
}
