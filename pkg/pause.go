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

func (self *Pause) Eval(scope *Scope, stack *Slice) error {
	t := scope.thread
	
	if t == nil {
		return scope.Error(self.form.Pos(), "No active thread")
	}

	var result Slice
	
	if err := t.scope.EvalOps(self.resultOps, &result); err != nil {
		return err
	}

	t.Pause(result.items)
	return nil
}
