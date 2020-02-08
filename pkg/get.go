package gfoo

type Get struct {
	OpBase
	key string
}

func NewGet(src Form, key string) *Get {
	o := new(Get)
	o.OpBase.Init(src)
	o.key = key
	return o
}

func (self *Get) Evaluate(gfoo *GFoo, scope *Scope) error {
	b := scope.Get(self.key)

	if b == nil {
		return gfoo.Errorf(self.source.Position(), "Unknown identifier: %v", self.key)
	}
	
	gfoo.Push(b.value)
	return nil
}

