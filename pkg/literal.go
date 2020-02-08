package gfoo

type Literal struct {
	FormBase
	val Val
}

func NewLiteral(pos Pos, dataType Type, data interface{}) *Literal {
	f := new(Literal)
	f.FormBase.Init(pos)
	f.val.Init(dataType, data)
	return f
}

func (self *Literal) Compile(gfoo *GFoo, scope *Scope, out []Op) ([]Op, error) {
	return append(out, NewPush(self, self.val)), nil
}

func (self *Literal) Quote() Val {
	return self.val
}
