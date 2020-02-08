package gfoo

type Bindings = map[string]Binding

type Scope struct {
	bindings Bindings
}

func (self *Scope) Init() {
	self.bindings = make(Bindings)
}

func (self *Scope) Get(key string) *Value {
	if found, ok := self.bindings[key]; ok {
		return &found.value
	}

	return nil
}

func (self *Scope) Set(key string, dataType Type, data interface{}) {
	self.bindings[key] = NewBinding(dataType, data)
}
