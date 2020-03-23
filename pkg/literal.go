package gfoo

import (
	"io"
)

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

func (self *Literal) Dump(out io.Writer) error {
	return self.val.Dump(out)
}

func (self *Literal) Quote(scope *Scope, thread *Thread, registers *Slice, pos Pos) (Val, error) {
	return self.val, nil
}
