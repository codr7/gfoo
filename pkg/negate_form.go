package gfoo

import (
	"io"
)

type NegateForm struct {
	FormBase
	form Form
}

func NewNegateForm(form Form, pos Pos) *NegateForm {
	return new(NegateForm).Init(form, pos)
}

func (self *NegateForm) Init(form Form, pos Pos) *NegateForm {
	self.FormBase.Init(pos)
	self.form = form
	return self
}

func (self *NegateForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	var err error
	
	if out, err = self.form.Compile(in, out, scope); err != nil {
		return out, err
	}
	
	return append(out, NewNegate(self)), nil
}

func (self *NegateForm) Do(action func(Form) error) error {
	return self.form.Do(action)
}

func (self *NegateForm) Dump(out io.Writer) error {	
	if _, err := io.WriteString(out, "!"); err != nil {
		return err
	}

	return self.form.Dump(out)
}

func (self *NegateForm) Quote(in *Forms, scope *Scope, thread *Thread, registers *Slice, pos Pos) (Val, error) {
	v, err := self.form.Quote(in, scope, thread, registers, pos)

	if err != nil {
		return Nil, err
	}

	v.Negate()
	return v, nil
}
