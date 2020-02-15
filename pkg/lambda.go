package gfoo

type Lambda struct {
	body []Op
	scope *Scope
}

func NewLambda(body []Op, scope *Scope) *Lambda {
	return new(Lambda).Init(body, scope)
}

func (self *Lambda) Init(body []Op, scope *Scope) *Lambda {
	self.body = body
	self.scope = scope
	return self
}

func (self *Lambda) Call(stack *Slice) error {
	return self.scope.Clone().Evaluate(self.body, stack)
}
