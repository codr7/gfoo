package gfoo

type Pause struct {
	OpBase
	result []Op
}

func NewPause(form Form, result []Op) *Pause {
	p := new(Pause)
	p.OpBase.Init(form)
	p.result = result
	return p
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
