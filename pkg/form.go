package gfoo

type Form interface {
	Dumper
	Compile(in *Forms, out []Op, scope *Scope) ([]Op, error)
	Do(action func(Form) error) error
	Pos() Pos
	Quote(in *Forms, scope *Scope, thread *Thread, registers []Val, pos Pos) (Val, error)
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
