package gfoo

type Capture struct {
	OpBase
	scope *Scope
}

func NewCapture(form Form, scope *Scope) *Capture {
	op := new(Capture)
	op.OpBase.Init(form)
	op.scope = scope
	return op
}

func (self *Capture) Eval(scope *Scope, stack *Slice) error {
	self.scope.Extend(scope)
	return nil
}

