package gfoo

type Binding struct {
	val Val
	scope *Scope
}

func NewBinding(scope *Scope, val Val) Binding {
	var b Binding
	b.Init(scope, val)
	return b
}

func (self *Binding) Init(scope *Scope, val Val) {
	self.scope = scope
	self.val = val
}
