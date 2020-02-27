package gfoo

type Let struct {
	OpBase
	key string
}

func NewLet(form Form, key string) *Let {
	op := new(Let)
	op.OpBase.Init(form)
	op.key = key
	return op
}

func (self *Let) Evaluate(scope *Scope, stack *Slice) error {	
	p := self.form.Pos()
	v, ok := stack.Pop()

	if !ok {
		return scope.Error(p, "Missing value: %v", self.key)
	}

	scope.Set(self.key, v)
	return nil
}

