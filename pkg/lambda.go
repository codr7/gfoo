package gfoo

type Lambda struct {
	argCount int
	body []Op
	registers Registers
}

func NewLambda(argCount int, body []Op, registers []Val) *Lambda {
	return new(Lambda).Init(argCount, body, registers)
}

func (self *Lambda) Init(argCount int, body []Op, registers []Val) *Lambda {
	self.body = body
	copy(self.registers[:], registers)
	return self
}

func (self *Lambda) Call(thread *Thread, stack *Stack, pos Pos) error {
	if sl := stack.Len(); sl < self.argCount {
		return Error(pos, "Not enough arguments: %v (%v)", sl, self.argCount)
	}

	return EvalOps(self.body, thread, self.registers[:], stack)
}
