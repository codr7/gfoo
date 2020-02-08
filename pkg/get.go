package gfoo

type Get struct {
	OpBase
	key string
}

func NewGet(form Form, key string) *Get {
	o := new(Get)
	o.OpBase.Init(form)
	o.key = key
	return o
}

func (self *Get) Evaluate(gfoo *GFoo, scope *Scope) error {
	found := scope.Get(self.key)

	if found == nil || found.val.data == nil {
		return gfoo.Error(self.form.Pos(), "Unknown identifier: %v", self.key)
	}

	v := &found.val
	gfoo.Push(v.dataType, v.data)
	return nil
}

