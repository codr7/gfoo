package gfoo

type Binding struct {
	val Val
	scope *Scope
}

func NewBinding(scope *Scope, dataType Type, data interface{}) Binding {
	var b Binding
	b.scope = scope
	b.val.Init(dataType, data)
	return b
}
