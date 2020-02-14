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

func (self *ThreadOp) Evaluate(stack *Slice, scope *Scope) error {
	t := NewThread(self.body, scope)
	
	if err := scope.Evaluate(self.args, &t.stack); err != nil {
		return err
	}

	t.Start()
	stack.Push(NewVal(&TThread, t))
	return nil
}

