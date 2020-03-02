package gfoo

type ScopeDef struct {
	OpBase
	keys []string
	values []Op
}

func NewScopeDef(form Form, keys []string, values []Op) *ScopeDef {
	op := new(ScopeDef)
	op.OpBase.Init(form)
	op.keys = keys
	op.values = values
	return op
}

func (self *ScopeDef) Eval(scope *Scope, stack *Slice) error {
	var values Slice
	
	if err := scope.EvalOps(self.values, &values); err != nil {
		return err
	}

	out := NewScope()

	for i := len(self.keys)-1; i >= 0; i-- {
		k := self.keys[i]
		v := values.Pop()

		if v == nil {
			return scope.Error(self.form.Pos(), "Missing value: %v", k)
		}
		
		out.Set(k, *v)
	}

	stack.Push(NewVal(&TScope, out))
	return nil
}
