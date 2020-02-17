package gfoo

type Reset struct {
	OpBase
}

func NewReset(form Form) *Reset {
	op := new(Reset)
	op.OpBase.Init(form)
	return op
}

func (self *Reset) Evaluate(scope *Scope, stack *Slice) error {
	stack.Reset()
	return nil
}
