package gfoo

type MethodOp struct {
	OpBase
	method *Method
}

func NewMethodOp(form Form, method *Method) *MethodOp {
	op := new(MethodOp)
	op.OpBase.Init(form)
	op.method = method
	return op
}

func (self *MethodOp) Eval(thread *Thread, registers, stack *Stack) error {
	self.method.registers.items = append(self.method.registers.items, registers.items...)
	return nil
}
