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
	key := self.key
	var source *Val
	
	if key[0] == '.' {
		key = key[1:]

		if source = stack.Pop(); source == nil {
			return scope.Error(self.form.Pos(), "Missing source: %v", self.key)
		}
	} else {
		s := NewVal(&TScope, scope)
		source = &s
	}

	v, err := source.Get(key, scope, self.form.Pos())
	
	if err != nil {
		return err
	}
	
	stack.Push(v)
	return nil
}

