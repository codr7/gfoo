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

func (self *MethodOp) Eval(thread *Thread, registers []Val, stack *Stack) error {
	copy(self.method.registers[:], registers)
	return nil
}
