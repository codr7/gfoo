package gfoo

type Push struct {
	OpBase
	dataType Type
	data interface{}
}

func NewPush(form Form, dataType Type, data interface{}) *Push {
	o := new(Push)
	o.OpBase.Init(form)
	o.dataType = dataType
	o.data = data
	return o
}

func (self *Push) Evaluate(gfoo *GFoo, scope *Scope) error {
	gfoo.Push(self.dataType, self.data)
	return nil
}
