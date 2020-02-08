package gfoo

type Binding struct {
	val Val
}

func NewBinding(dataType Type, data interface{}) Binding {
	var b Binding
	b.val.Init(dataType, data)
	return b
}
