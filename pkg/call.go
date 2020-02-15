package gfoo

type Call struct {
	OpBase
	target *Val
	args []Op
}

func NewCall(form Form, target *Val, args []Op) *Call {
	o := new(Call)
	o.OpBase.Init(form)
	o.target = target
	o.args = args
	return o
}

func (self *Call) Evaluate(scope *Scope, stack *Slice) error {
	t := self.target
	
	if t == nil {
		if v, ok := stack.Pop(); ok {
			t = &v
		}
	}

	if t == nil {
		scope.Error(self.form.Pos(), "Missing call target")
	}

	if err := scope.Evaluate(self.args, stack); err != nil {
		return err
	}
	
	return t.Call(scope, stack, self.form.Pos())
}
