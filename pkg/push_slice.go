package gfoo

type PushSlice struct {
	OpBase
	ops []Op
}

func NewPushSlice(form Form, ops []Op) *PushSlice {
	o := new(PushSlice)
	o.OpBase.Init(form)
	o.ops = ops
	return o
}

func (self *PushSlice) Evaluate(gfoo *GFoo, scope *Scope) error {
	i := gfoo.stack.Len()

	if err := gfoo.Evaluate(self.ops, scope); err != nil {
		return err
	}

	gfoo.Push(&TSlice, NewSlice(gfoo.stack.Cut(i)))
	return nil
}

