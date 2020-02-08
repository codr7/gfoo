package gfoo

type Drop struct {
	OpBase
}

func NewDrop(form Form) *Drop {
	o := new(Drop)
	o.OpBase.Init(form)
	return o
}

func (self *Drop) Evaluate(gfoo *GFoo, scope *Scope) error {
	if v := gfoo.Pop(); v == nil {
		return gfoo.Error(self.form.Pos(), "Nothing to drop")
	}
	
	return nil
}
