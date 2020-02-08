package gfoo

type Push struct {
	OpBase
	val Val
}

func NewPush(form Form, val Val) *Push {
	o := new(Push)
	o.OpBase.Init(form)
	o.val = val
	return o
}

func (self *Push) Evaluate(gfoo *GFoo, scope *Scope) error {
	gfoo.Push(self.val)
	return nil
}

