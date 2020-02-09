package gfoo

type Reset struct {
	OpBase
}

func NewReset(form Form) *Reset {
	o := new(Reset)
	o.OpBase.Init(form)
	return o
}

func (self *Reset) Evaluate(gfoo *GFoo, scope *Scope) error {
	gfoo.stack.Reset()
	return nil
}
