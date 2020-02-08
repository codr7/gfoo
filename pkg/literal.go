package gfoo

type Literal struct {
	FormBase
	dataType Type
	data interface{}
}

func NewLiteral(pos Pos, dataType Type, data interface{}) *Literal {
	f := new(Literal)
	f.FormBase.Init(pos)
	f.dataType = dataType
	f.data = data
	return f
}

func (self *Literal) Compile(gfoo *GFoo, scope *Scope, in *Forms, out []Op) ([]Op, error) {
	return append(out, NewPush(self, self.dataType, self.data)), nil
}

func (self *Literal) Quote() Val {
	return NewVal(self.dataType, self.data)
}
