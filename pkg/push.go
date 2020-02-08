package gfoo

type Push struct {
	OpBase
	val Val
}

func NewPush(src Form, val Val) *Push {
	o := new(Push)
	o.OpBase.Init(src)
	o.val = val
	return o
}

func (self *Push) Evaluate(gfoo *GFoo, scope *Scope) error {
	gfoo.Push(self.val)
	return nil
}

