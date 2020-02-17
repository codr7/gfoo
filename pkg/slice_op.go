package gfoo

type SliceOp struct {
	OpBase
	body []Op
}

func NewSliceOp(form Form, body []Op) *SliceOp {
	op := new(SliceOp)
	op.OpBase.Init(form)
	op.body = body
	return op
}

func (self *SliceOp) Evaluate(scope *Scope, stack *Slice) error {
	s := NewSlice(nil)
	
	if err := scope.Evaluate(self.body, s); err != nil {
		return err
	}

	stack.Push(NewVal(&TSlice, s))
	return nil
}

