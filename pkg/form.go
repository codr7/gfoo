package gfoo

import (
	"io"
	"strings"
)

type Form interface {
	Compile(in *Forms, out []Op, scope *Scope) ([]Op, error)
	Do(action func(Form) error) error
	Dump(out io.Writer) error
	Pos() Pos
	Quote(scope *Scope, pos Pos) (Val, error)
}

type FormBase struct {
	pos Pos
}

func (self *FormBase) Init(pos Pos) {
	self.pos = pos
}

func (self *FormBase) Pos() Pos {
	return self.pos
}

func FormString(form Form) string {
	var out strings.Builder
	form.Dump(&out);
	return out.String()
}
