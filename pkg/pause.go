package gfoo

type Pause struct {
	OpBase
	result []Op
}

func NewPause(form Form, result []Op) *Pause {
	op := new(Pause)
	op.OpBase.Init(form)
	op.result = result
	return op
}

func (self *Pause) Evaluate(scope *Scope, stack *Slice) error {
	t := scope.thread
	
	if t == nil {
		return scope.Error(self.form.Pos(), "No active thread")
	}

	if err := t.scope.Evaluate(self.result, &t.result); err != nil {
		return err
	}

	t.Pause()
	return nil
}
