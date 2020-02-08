package gfoo

type Let struct {
	OpBase
	key string
}

func NewLet(form Form, key string) *Let {
	o := new(Let)
	o.OpBase.Init(form)
	o.key = key
	return o
}

func (self *Let) Evaluate(gfoo *GFoo, scope *Scope) error {
	p := self.form.Pos()
	v := gfoo.Pop()

	if v == nil {
		return gfoo.Error(p, "Missing value: %v", self.key)
	}
	
	scope.Set(self.key, v.dataType, v.data)
	return nil
}

