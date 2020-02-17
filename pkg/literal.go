package gfoo

type Literal struct {
	FormBase
	val Val
}

func NewLiteral(val Val, pos Pos) *Literal {
	f := new(Literal)
	f.FormBase.Init(pos)
	f.val = val
	return f
}

func (self *Literal) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewPush(self, self.val)), nil
}

func (self *Literal) Do(action func(Form) error) error {
	return action(self)
}

func (self *Literal) Quote(scope *Scope) (Val, error) {
	return self.val, nil
}
