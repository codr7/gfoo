package gfoo

import (
	"io"
)

type PairForm struct {
	FormBase
	left, right Form
}

func NewPairForm(left, right Form, pos Pos) *PairForm {
	return new(PairForm).Init(left, right, pos)
}

func (self *PairForm) Init(left, right Form, pos Pos) *PairForm {
	self.FormBase.Init(pos)
	self.left = left
	self.right = right
	return self
}

func (self *PairForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	var left, right []Op
	var err error
	
	if left, err = self.left.Compile(nil, nil, scope); err != nil {
		return out, nil
	}

	if right, err = self.right.Compile(nil, nil, scope); err != nil {
		return out, nil
	}

	return append(out, NewPairOp(self, left, right)), nil
}

func (self *PairForm) Do(action func(Form) error) error {
	if err := self.left.Do(action); err != nil {
		return err
	}

	if err := self.right.Do(action); err != nil {
		return err
	}

	return nil
}

func (self *PairForm) Dump(out io.Writer) error {
	if _, err := io.WriteString(out, ","); err != nil {
		return err
	}

	if err := self.left.Dump(out); err != nil {
		return err
	}

	if _, err := io.WriteString(out, " "); err != nil {
		return err
	}
	
	if err := self.right.Dump(out); err != nil {
		return err
	}

	return nil
}

func (self *PairForm) Quote(in *Forms, scope *Scope, thread *Thread, registers []Val, pos Pos) (Val, error) {
	var left, right Val
	var err error

	if left, err = self.left.Quote(nil, scope, thread, registers, pos); err != nil {
		return Nil, err
	}

	if right, err = self.right.Quote(nil, scope, thread, registers, pos); err != nil {
		return Nil, err
	}

	return NewVal(&TPair, NewPair(left, right)), nil
}

