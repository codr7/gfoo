package gfoo

type PushSlice struct {
	OpBase
	ops []Op
}

func NewPushSlice(src Form, ops []Op) *PushSlice {
	o := new(PushSlice)
	o.OpBase.Init(src)
	o.ops = ops
	return o
}

func (self *PushSlice) Evaluate(gfoo *GFoo, scope *Scope) error {
	i := len(gfoo.stack)

	if err := gfoo.Evaluate(self.ops, scope); err != nil {
		return err
	}
	
	gfoo.Push(NewValue(&Slice, gfoo.stack[i:]))
	gfoo.stack = gfoo.stack[:i]
	return nil
}

