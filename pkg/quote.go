package gfoo

import (
	"io"
)

type Quote struct {
	FormBase
	form Form
}

func NewQuote(form Form, pos Pos) *Quote {
	return new(Quote).Init(form, pos)
}

func (self *Quote) Init(form Form, pos Pos) *Quote {
	self.FormBase.Init(pos)
	self.form = form
	return self
}

func (self *Quote) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return append(out, NewQuoteOp(self.form, scope)), nil
}

func (self *Quote) Do(action func(Form) error) error {
	return self.form.Do(action)
}

func (self *Quote) Dump(out io.Writer) error {
	if _, err := io.WriteString(out, "'"); err != nil {
		return err
	}
	
	return self.form.Dump(out)
}

func (self *Quote) Quote(in *Forms, scope *Scope, thread *Thread, registers *Slice, pos Pos) (Val, error) {
	v, err := self.form.Quote(in, scope, thread, registers, pos)

	if err != nil {
		return Nil, err
	}
	
	return NewVal(&TQuote, v), nil
}
