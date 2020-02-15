package gfoo

type Reset struct {
	OpBase
}

func NewReset(form Form) *Reset {
	o := new(Reset)
	o.OpBase.Init(form)
	return o
}

func (self *Reset) Evaluate(scope *Scope, stack *Slice) error {
	stack.Reset()
	return nil
}
