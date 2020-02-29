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

func (self *Branch) Evaluate(scope *Scope, stack *Slice) error {
	v := stack.Pop()
	
	if v == nil {
		scope.Error(self.form.Pos(), "Missing condition")
	}

	var body []Op
	
	if v.Bool() {
		body = self.trueBody
	} else {
		body = self.falseBody
	}
	
	return scope.Evaluate(body, stack)
}
