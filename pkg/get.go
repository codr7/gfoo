package gfoo

type Get struct {
	OpBase
	key string
}

func NewGet(form Form, key string) *Get {
	o := new(Get)
	o.OpBase.Init(form)
	o.key = key
	return o
}

func (self *Get) Evaluate(stack *Slice, scope *Scope) error {
	found := scope.Get(self.key)

	if found == nil || found.val == NilVal {
		return scope.Error(self.form.Pos(), "Unknown identifier: %v", self.key)
	}

	stack.Push(found.val)
	return nil
}

