package gfoo

type Call struct {
	OpBase
	target *Val
	args []Op
}

func NewCall(form Form, target *Val, args []Op) *Call {
	op := new(Call)
	op.OpBase.Init(form)
	op.target = target
	op.args = args
	return op
}

func (self *Call) Evaluate(scope *Scope, stack *Slice) error {
	t := self.target
	
	if t == nil {
		if t = stack.Pop(); t == nil {
			scope.Error(self.form.Pos(), "Missing target")
		}
	}

	if err := scope.Evaluate(self.args, stack); err != nil {
		return err
	}
	
	return t.Call(scope, stack, self.form.Pos())
}
