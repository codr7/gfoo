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

func (self *Push) Evaluate(vm *VM, stack *Slice, scope *Scope) error {
	stack.Push(self.dataType, self.data)
	return nil
}
