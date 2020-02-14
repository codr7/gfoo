package gfoo

type Op interface {
	Evaluate(stack *Slice, scope *Scope) error
}

type OpBase struct {
	form Form
}

func (self *OpBase) Init(form Form) {
	self.form = form
}
