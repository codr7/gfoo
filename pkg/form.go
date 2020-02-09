package gfoo

type Form interface {
	Compile(vm *VM, scope *Scope, in *Forms, out []Op) ([]Op, error)
	Pos() Pos
	Quote() Val
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
