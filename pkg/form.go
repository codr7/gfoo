package gfoo

type Form interface {
	Compile(in *Forms, out []Op, scope *Scope) ([]Op, error)
	Pos() Pos
	Quote(scope *Scope) (Val, error)
}

type FormBase struct {
	pos Pos
}

func (self *FormBase) Init(pos Pos) {
	self.pos = pos
}

func (self *FormBase) Pos() Pos {
	return self.pos
}
