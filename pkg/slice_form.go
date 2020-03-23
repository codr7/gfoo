package gfoo

import (
	"io"
)

type SliceForm struct {
	FormBase
	body []Form
}

func NewSliceForm(body []Form, pos Pos) *SliceForm {
	return new(SliceForm).Init(body, pos)
}

func (self *SliceForm) Init(body []Form, pos Pos) *SliceForm {
	self.FormBase.Init(pos)
	self.body = body
	return self
}

func (self *SliceForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	ops, err := scope.Compile(self.body, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewSliceOp(self, ops)), nil
}

func (self *SliceForm) Do(action func(Form) error) error {
	for _, f := range self.body {
		if err := f.Do(action); err != nil {
			return err
		}
	}

	return nil
}

func (self *SliceForm) Dump(out io.Writer) error {
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

func (self *SliceForm) Quote(scope *Scope, thread *Thread, registers *Slice, pos Pos) (Val, error) {
	ops, err := NewScope(scope).Compile(self.body, nil)

	if err != nil {
		return Nil, err
	}

	v := NewSlice(nil)
	
	if err = EvalOps(ops, nil, NewSlice(nil), v); err != nil {
		return Nil, err
	}
		
	return NewVal(&TSlice, v), nil
}
