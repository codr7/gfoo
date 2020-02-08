package gfoo

type Binding struct {
	value Value
}

func NewBinding(dataType Type, data interface{}) Binding {
	var b Binding
	b.value.Init(dataType, data)
	return b
}
