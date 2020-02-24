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

func (self *Check) Evaluate(scope *Scope, stack *Slice) error {
	condStack := stack.Clone()
	
	if err := scope.Evaluate(self.condOps, stack); err != nil {
		return err
	}

	result, ok := stack.Pop()
	
	if !ok {
		return scope.Error(self.form.Pos(), "Missing result")
	}

	if !result.Bool() {
		return scope.Error(self.form.Pos(), "Check failed: %v %v", self.cond, condStack)	
	}
	
	return nil
}
