package gfoo

type Push struct {
	OpBase
	value Value
}

func NewPush(src Form, val Value) *Push {
	o := new(Push)
	o.OpBase.Init(src)
	o.value = val
	return o
}

func (self *Push) Evaluate(gfoo *GFoo, scope *Scope) error {
	gfoo.Push(self.value)
	return nil
}

