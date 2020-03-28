
package gfoo

import (
	"io"
)

type Unquote struct {
	FormBase
	form Form
}

func NewUnquote(form Form, pos Pos) *Unquote {
	return new(Unquote).Init(form, pos)
}

func (self *Unquote) Init(form Form, pos Pos) *Unquote {
	self.FormBase.Init(pos)
	self.form = form
	return self
}

func (self *Unquote) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	return nil, Error(self.pos, "Unquote outside of quoted context")
}

func (self *Unquote) Do(action func(Form) error) error {
	return self.form.Do(action)
}

func (self *Unquote) Dump(out io.Writer) error {
	return self.form.Dump(out)
}

func (self *Unquote) Quote(in *Forms, scope *Scope, thread *Thread, registers *Stack, pos Pos) (Val, error) {
	ops, err := self.form.Compile(in, nil, scope)

	if err != nil {
		return Nil, err
	}

	var stack Stack
	stack.Init(nil)

	if err = EvalOps(ops, thread, registers, &stack); err != nil {
		return Nil, err
	}

	v := stack.Pop()

	if v == nil {
		return Nil, Error(pos, "Empty unquote")
	}
	
	return *v, nil
}
