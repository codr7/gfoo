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
	
	if key[0] == '.' {
		sv, ok := stack.Pop();
		
		if !ok {
			return scope.Error(self.form.Pos(), "Missing scope: %v", self.key)
		}

		if sv.dataType != &TScope {
			return scope.Error(self.form.Pos(), "Expected scope: %v", sv)
		}
		
		scope = sv.data.(*Scope)
		key = key[1:]
	}
	
	found := scope.Get(key)

	if found == nil || found.val == Nil {
		return scope.Error(self.form.Pos(), "Unknown identifier: %v", key)
	}

	stack.Push(found.val)
	return nil
}

