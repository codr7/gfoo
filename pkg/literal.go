package gfoo

type Literal struct {
	value Value
}

func NewLiteral(dataType Type, data interface{}) *Literal {
	lit := new(Literal)
	lit.value.Init(dataType, data)
	return lit
}

func (lit *Literal) Quote() Value {
	return lit.value
}
