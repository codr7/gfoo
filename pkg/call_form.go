package gfoo

import (
	"io"
)

type CallForm struct {
	FormBase
	id *Id
	args *Group
}

func NewCallForm(id *Id, args *Group, pos Pos) *CallForm {
	return new(CallForm).Init(id, args, pos)
}

func (self *CallForm) Init(id *Id, args *Group, pos Pos) *CallForm {
	self.FormBase.Init(pos)
	self.id = id
	self.args = args
	return self
}

func (self *CallForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	in.Push(self.args);
	return self.id.Compile(in, out, scope)
}

func (self *CallForm) Do(action func(Form) error) error {
	if err := self.id.Do(action); err != nil {
		return err
	}
	
	return self.args.Do(action)
}

func (self *CallForm) Dump(out io.Writer) error {
	if err := self.id.Dump(out); err != nil {
		return err
	}
	
	return self.args.Dump(out)
}

func (self *CallForm) Quote(in *Forms, scope *Scope, thread *Thread, registers *Stack, pos Pos) (Val, error) {
	out := make([]Val, len(self.args.body)+1)
	var err error

	if out[0], err = self.id.Quote(nil, scope, thread, registers, pos); err != nil {
		return Nil, err
	}
	
	for i, f := range self.args.body {
		if out[i+1], err = f.Quote(nil, scope, thread, registers, pos); err != nil {
			return Nil, err
		}
	}
	
	return NewVal(&TCall, out), nil
}
