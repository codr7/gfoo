package gfoo

type ThreadOp struct {
	OpBase
	args, body []Op
}

func NewThreadOp(form Form, args []Op, body []Op) *ThreadOp {
	op := new(ThreadOp)
	op.OpBase.Init(form)
	op.args = args
	op.body = body
	return op
}

func (self *ThreadOp) Eval(thread *Thread, registers []Val, stack *Stack) error {
	t := NewThread(self.body)
	
	if err := EvalOps(self.args, thread, registers, &t.stack); err != nil {
		return err
	}

	t.Start()
	stack.Push(NewVal(&TThread, t))
	return nil
}

