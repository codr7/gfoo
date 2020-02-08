package gfoo

type Op interface {
	Evaluate(gfoo *GFoo, scope *Scope) error
}

type OpBase struct {
	form Form
}

func (self *OpBase) Init(form Form) {
	self.form = form
}
