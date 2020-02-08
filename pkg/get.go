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
	b := scope.Get(self.key)

	if b == nil {
		return gfoo.Error(self.form.Pos(), "Unknown identifier: %v", self.key)
	}
	
	gfoo.Push(b.val)
	return nil
}

