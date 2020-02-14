package gfoo

type Lambda struct {
	forms []Form
	ops []Op
	scope *Scope
}

func NewLambda(forms []Form, ops []Op, scope *Scope) *Lambda {
	return new(Lambda).Init(forms, ops, scope)
}

func (self *Lambda) Init(forms []Form, ops []Op, scope *Scope) *Lambda {
	self.forms = forms
	self.ops = ops
	self.scope = scope
	return self
}

func (self *Lambda) Call(stack *Slice) error {
	return self.scope.Clone(true).Evaluate(self.ops, stack)
}
