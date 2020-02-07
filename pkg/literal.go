package gfoo

type Literal struct {
	value *Value
}

func NewLiteral(val *Value) *Literal {
	return &Literal{value: val}
}

func (lit *Literal) Quote() *Value {
	return lit.value
}
