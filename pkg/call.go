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

func (self *Call) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	t := self.target
	
	if t == nil {
		t = stack.Pop()
	}

	if t == nil {
		vm.Error(self.form.Pos(), "Missing call target")
	}
	
	return t.Call(vm, stack)
}
