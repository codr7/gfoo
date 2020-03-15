package gfoo

type PeekVal struct {
	OpBase
}

func NewPeekVal(form Form) *PeekVal {
	op := new(PeekVal)
	op.OpBase.Init(form)
	return op
}

func (self *PeekVal) Eval(scope *Scope, stack *Slice) error {
	if v := scope.val.Peek(); v == nil {
		stack.Push(Nil)
	} else {
		stack.Push(*v)
	}

	return nil
}

