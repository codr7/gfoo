package gfoo

type Drop struct {
	OpBase
}

func NewDrop(form Form) *Drop {
	op := new(Drop)
	op.OpBase.Init(form)
	return op
}

func (self *Drop) Evaluate(scope *Scope, stack *Slice) error {
	if _, ok := stack.Pop(); !ok {
		return scope.Error(self.form.Pos(), "Nothing to drop")
	}
	
	return nil
}
