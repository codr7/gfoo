package gfoo

type Dup struct {
	OpBase
}

func NewDup(form Form) *Dup {
	op := new(Dup)
	op.OpBase.Init(form)
	return op
}

func (self *Dup) Evaluate(scope *Scope, stack *Slice) error {
	v := stack.Peek()
	
	if v == nil {
		return scope.Error(self.form.Pos(), "Nothing to dup")
	}

	stack.Push(*v)
	return nil
}
