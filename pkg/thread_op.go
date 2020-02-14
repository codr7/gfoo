package gfoo

type ThreadOp struct {
	OpBase
	args, body []Op
}

func NewThreadOp(form Form, args []Op, body []Op) *ThreadOp {
	o := new(ThreadOp)
	o.OpBase.Init(form)
	o.args = args
	o.body = body
	return o
}

func (self *ThreadOp) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	t := NewThread(self.body, scope)
	
	if err := scope.vm.Evaluate(self.args, &t.stack, scope); err != nil {
		return err
	}

	t.Start()
	stack.Push(NewVal(&TThread, t))
	return nil
}

