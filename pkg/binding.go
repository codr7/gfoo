package gfoo

type Binding struct {
	val Val
	scope *Scope
}

func NewBinding(scope *Scope, val Val) Binding {
	var b Binding
	b.scope = scope
	b.val = val
	return b
}
