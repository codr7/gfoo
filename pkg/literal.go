package gfoo

type Literal struct {
	FormBase
	val Val
}

func NewLiteral(pos Pos, val Val) *Literal {
	f := new(Literal)
	f.FormBase.Init(pos)
	f.val = val
	return f
}

func (self *Literal) Compile(in *Forms, out []Op, vm *VM, scope *Scope) ([]Op, error) {
	return append(out, NewPush(self, self.val)), nil
}

func (self *Literal) Quote(vm *VM, scope *Scope) (Val, error) {
	return self.val, nil
}
