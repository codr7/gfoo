package gfoo

import (
	"io"
)

type StackForm struct {
	FormBase
	body []Form
}

func NewStackForm(body []Form, pos Pos) *StackForm {
	return new(StackForm).Init(body, pos)
}

func (self *StackForm) Init(body []Form, pos Pos) *StackForm {
	self.FormBase.Init(pos)
	self.body = body
	return self
}

func (self *StackForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	ops, err := scope.Compile(self.body, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewStackOp(self, ops)), nil
}

func (self *StackForm) Do(action func(Form) error) error {
	for _, f := range self.body {
		if err := f.Do(action); err != nil {
			return err
		}
	}

	return nil
}

func (self *StackForm) Dump(out io.Writer) error {
 	io.WriteString(out, "[")
	
	for i, f := range self.body {
		if i > 0 {
			io.WriteString(out, " ")
		}

		if err := f.Dump(out); err != nil {
			return err
		}
	}

	io.WriteString(out, "]")
	return nil
}

func (self *StackForm) Quote(in *Forms, scope *Scope, thread *Thread, registers *Stack, pos Pos) (Val, error) {
	out := make([]Val, len(self.body))
	var err error
	
	for i, f := range self.body {
		if out[i], err = f.Quote(nil, scope, thread, registers, pos); err != nil {
			return Nil, err
		}
	}
		
	return NewVal(&TStack, NewStack(out)), nil
}
