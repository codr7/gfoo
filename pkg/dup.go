package gfoo

type Dup struct {
	OpBase
}

func NewDup(form Form) *Dup {
	o := new(Dup)
	o.OpBase.Init(form)
	return o
}

func (self *Dup) Evaluate(gfoo *GFoo, scope *Scope) error {
	v := gfoo.Peek()
	
	if v == nil {
		return gfoo.Error(self.form.Pos(), "Nothing to dup")
	}

	gfoo.Push(v.dataType, v.data)
	return nil
}
