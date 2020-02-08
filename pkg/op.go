package gfoo

type Op interface {
	Evaluate(gfoo *GFoo, scope *Scope) error
}

type OpBase struct {
	source Form
}

func (self *OpBase) Init(src Form) {
	self.source = src
}
