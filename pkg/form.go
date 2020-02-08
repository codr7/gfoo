package gfoo

type Form interface {
	Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error)
	Position() Position
	Quote() Value
}

type FormBase struct {
	position Position
}

func (self *FormBase) Init(pos Position) {
	self.position = pos
}

func (self *FormBase) Position() Position {
	return self.position
}
