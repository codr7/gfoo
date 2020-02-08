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
	i := len(gfoo.stack)

	if err := gfoo.Evaluate(self.ops, scope); err != nil {
		return err
	}

	n := len(gfoo.stack) - i
	items := make([]Val, n)
	copy(items, gfoo.stack[i:])
	gfoo.stack = gfoo.stack[:i]
	gfoo.Push(NewVal(&TSlice, items))
	return nil
}

