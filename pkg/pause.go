package gfoo

type Pause struct {
	OpBase
	resultOps []Op
}

func NewPause(form Form, resultOps []Op) *Pause {
	op := new(Pause)
	op.OpBase.Init(form)
	op.resultOps = resultOps
	return op
}

func (self *Pause) Eval(thread *Thread, registers, stack *Stack) error {
	if thread == nil {
		return Error(self.form.Pos(), "No active thread")
	}

	var result Stack
	
	if err := EvalOps(self.resultOps, thread, registers, &result); err != nil {
		return err
	}

	thread.Pause(result.items)
	return nil
}
