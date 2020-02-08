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

func (self *Id) Compile(gfoo *GFoo, scope *Scope, in *Forms, out []Op) ([]Op, error) {
	if b := scope.Get(self.name); b != nil {
		v := &b.val
		
		if v.dataType == &TMacro {
			return v.data.(*Macro).Expand(gfoo, scope, self, in, out)
		}
		
		return append(out, NewPush(self, v.dataType, v.data)), nil
	}

	return append(out, NewGet(self, self.name)), nil
}

func (self *Id) Quote() Val {
	return NewVal(&TId, self.name)
}
