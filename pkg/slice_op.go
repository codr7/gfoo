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

func (self *SliceOp) Eval(thread *Thread, registers, stack *Slice) error {
	s := NewSlice(nil)
	
	if err := EvalOps(self.body, thread, registers, s); err != nil {
		return err
	}

	stack.Push(NewVal(&TSlice, s))
	return nil
}

