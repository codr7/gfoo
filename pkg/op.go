package gfoo

type Op interface {
	Eval(thread *Thread, registers, stack *Stack) error
}

type OpBase struct {
	form Form
}

func (self *OpBase) Init(form Form) {
	self.form = form
}

func EvalOps(ops []Op, thread *Thread, registers, stack *Stack) error {
	for _, op := range ops {
		if err := op.Eval(thread, registers, stack); err != nil {
			return err
		}
	}
	
	return nil
}
