package gfoo

type Literal struct {
	FormBase
	value Value
}

func NewLiteral(pos Position, dataType Type, data interface{}) *Literal {
	f := new(Literal)
	f.FormBase.Init(pos)
	f.value.Init(dataType, data)
	return f
}

func (self *Literal) Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error) {
	return append(out, NewPush(self, self.value)), nil
}

func (self *Literal) Quote() Value {
	return self.value
}
