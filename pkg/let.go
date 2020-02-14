package gfoo

type Let struct {
	OpBase
	key string
}

func NewLet(form Form, key string) *Let {
	o := new(Let)
	o.OpBase.Init(form)
	o.key = key
	return o
}

func (self *Let) Evaluate(stack *Slice, scope *Scope) error {
	p := self.form.Pos()
	v, ok := stack.Pop()

	if !ok {
		return scope.Error(p, "Missing value: %v", self.key)
	}
	
	scope.Set(self.key, v)
	return nil
}

