package gfoo

type Get struct {
	OpBase
	key string
}

func NewGet(form Form, key string) *Get {
	op := new(Get)
	op.OpBase.Init(form)
	op.key = key
	return op
}

func (self *Get) Eval(thread *Thread, registers, stack *Stack) error {
	var source *Val
	
	if source = stack.Pop(); source == nil {
		return Error(self.form.Pos(), "Missing source: %v", self.key)
	}
	
	v, err := source.Get(self.key, self.form.Pos())
	
	if err != nil {
		return err
	}
	
	stack.Push(v)
	return nil
}
