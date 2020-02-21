package gfoo

type Lambda struct {
	argCount int
	body []Op
	scope *Scope
}

func NewLambda(argCount int, body []Op, scope *Scope) *Lambda {
	return new(Lambda).Init(argCount, body, scope)
}

func (self *Lambda) Init(argCount int, body []Op, scope *Scope) *Lambda {
	self.body = body
	self.scope = scope
	return self
}

func (self *Lambda) Call(stack *Slice, pos Pos) error {
	if sl := stack.Len(); sl < self.argCount {
		self.scope.Error(pos, "Not enough arguments: %v (%v)", sl, self.argCount)
	}
	
	return self.scope.Clone().Evaluate(self.body, stack)
}
