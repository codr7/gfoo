package gfoo

type Bindings = map[string]Binding

type Scope struct {
	bindings Bindings
}

func (self *Scope) Init() *Scope {
	self.bindings = make(Bindings)
	return self
}

func (self *Scope) Clone() *Scope {
	out := new(Scope).Init()

	for k, b := range self.bindings {
		out.bindings[k] = b
	}
	
	return out
}

func (self *Scope) Get(key string) *Binding {
	if found, ok := self.bindings[key]; ok {
		return &found
	}

	return nil
}

func (self *Scope) Set(key string, val Val) {
	self.bindings[key] = NewBinding(self, val)
}
