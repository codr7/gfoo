package gfoo

type Op interface {
	Evaluate(scope *Scope, stack *Slice) error
}

type OpBase struct {
	form Form
}

func (self *OpBase) Init(form Form) {
	self.form = form
}
