package gfoo

type Check struct {
	OpBase
	cond Form
	condOps []Op
}

func NewCheck(form Form, cond Form, condOps []Op) *Check {
	op := new(Check)
	op.OpBase.Init(form)
	op.cond = cond
	op.condOps = condOps
	return op
}

func (self *Check) Eval(thread *Thread, registers, stack *Stack) error {
	condStack := stack.Clone()
	
	if err := EvalOps(self.condOps, thread, registers, stack); err != nil {
		return err
	}

	result := stack.Pop()
	
	if result == nil {
		return Error(self.form.Pos(), "Missing result")
	}

	if !result.Bool() {
		return Error(self.form.Pos(), "Check failed: %v %v", self.cond, condStack)	
	}
	
	return nil
}
