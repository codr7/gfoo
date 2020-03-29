package gfoo

type LambdaOp struct {
	OpBase
	method *Method
	argCount int
	body []Op
}

func NewLambdaOp(form Form, argCount int, body []Op) *LambdaOp {
	op := new(LambdaOp)
	op.OpBase.Init(form)
	op.argCount = argCount
	op.body = body
	return op
}

func (self *LambdaOp) Eval(thread *Thread, registers []Val, stack *Stack) error {
	stack.Push(NewVal(&TLambda, NewLambda(self.argCount, self.body, registers)))
	return nil
}
