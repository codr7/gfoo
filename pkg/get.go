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

func (self *Get) Evaluate(scope *Scope, stack *Slice) error {
	found := scope.Get(self.key)

	if found == nil || found.val == NilVal {
		return scope.Error(self.form.Pos(), "Unknown identifier: %v", self.key)
	}

	stack.Push(found.val)
	return nil
}

