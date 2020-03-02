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

func (self *Let) Eval(scope *Scope, stack *Slice) error {	
	v := stack.Pop()

	if v == nil {
		return scope.Error(self.form.Pos(), "Missing value: %v", self.key)
	}

	scope.Set(self.key, *v)
	return nil
}

