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

func (self *ThreadOp) Eval(scope *Scope, stack *Slice) error {
	t := NewThread(self.body, scope)
	
	if err := scope.EvalOps(self.args, &t.stack); err != nil {
		return err
	}

	t.Start()
	stack.Push(NewVal(&TThread, t))
	return nil
}

