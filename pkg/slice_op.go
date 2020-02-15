package gfoo

type SliceOp struct {
	OpBase
	ops []Op
}

func NewSliceOp(form Form, ops []Op) *SliceOp {
	o := new(SliceOp)
	o.OpBase.Init(form)
	o.ops = ops
	return o
}

func (self *SliceOp) Evaluate(scope *Scope, stack *Slice) error {
	s := NewSlice(nil)
	
	if err := scope.Evaluate(self.ops, s); err != nil {
		return err
	}

	stack.Push(NewVal(&TSlice, s))
	return nil
}

