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

	n := self.name
	
	if n[0] == '$' {
		n = scope.Unique(n)
	}
	
	return append(out, NewGet(self, n)), nil
}

func (self *Id) Quote(scope *Scope) (Val, error) {
	n := self.name
	
	if n[0] == '$' {
		n = scope.Unique(n)
	}

	return NewVal(&TId, n), nil
}
