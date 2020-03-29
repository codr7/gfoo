package gfoo

type Branch struct {
	OpBase
	trueBody, falseBody []Op
}

func NewBranch(form Form, trueBody, falseBody []Op) *Branch {
	op := new(Branch)
	op.OpBase.Init(form)
	op.trueBody = trueBody
	op.falseBody = falseBody
	return op
}

func (self *Branch) Eval(thread *Thread, registers []Val, stack *Stack) error {
	v := stack.Pop()
	
	if v == nil {
		Error(self.form.Pos(), "Missing condition")
	}

	var body []Op
	
	if v.Bool() {
		body = self.trueBody
	} else {
		body = self.falseBody
	}
	
	return EvalOps(body, thread, registers, stack)
}
