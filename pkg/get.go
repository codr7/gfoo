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

func (self *Get) Eval(scope *Scope, stack *Slice) error {
	key := self.key
	
	if key[0] == '.' {
		key = key[1:]
		var source *Val

		if source = stack.Pop(); source == nil {
			return scope.Error(self.form.Pos(), "Missing source: %v", self.key)
		}

		v, err := source.Get(key, scope, self.form.Pos())
		
		if err != nil {
			return err
		}
	
		stack.Push(v)
	} else {
		found := scope.Get(key)
		
		if found == nil || found.val == Undefined {
			return scope.Error(self.form.Pos(), "Unknown identifier: %v", key)
		}
		
		stack.Push(found.val)
	}

	return nil
}

