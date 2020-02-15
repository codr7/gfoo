package gfoo

type Call struct {
	OpBase
	target *Val
}

func NewCall(form Form, target *Val) *Call {
	o := new(Call)
	o.OpBase.Init(form)
	o.target = target
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
	
	return t.Call(scope, stack, self.form.Pos())
}
