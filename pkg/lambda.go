package gfoo

type Lambda struct {
	argCount int
	body []Op
	registers Stack
}

func NewLambda(argCount int, body []Op, registers *Stack) *Lambda {
	return new(Lambda).Init(argCount, body, registers)
}

func (self *Lambda) Init(argCount int, body []Op, registers *Stack) *Lambda {
	self.body = body
	self.registers.items = append(self.registers.items, registers.items...)
	return self
}

func (self *Lambda) Call(thread *Thread, stack *Stack, pos Pos) error {
	if sl := stack.Len(); sl < self.argCount {
		return Error(pos, "Not enough arguments: %v (%v)", sl, self.argCount)
	}

	return EvalOps(self.body, thread, &self.registers, stack)
}
