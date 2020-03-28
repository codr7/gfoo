package gfoo

import (
	//"fmt"
	"io"
)

type ScopeForm struct {
	FormBase
	body []Form
}

func NewScopeForm(body []Form, pos Pos) *ScopeForm {
	return new(ScopeForm).Init(body, pos)
}

func (self *ScopeForm) Init(body []Form, pos Pos) *ScopeForm {
	self.FormBase.Init(pos)
	self.body = body
	return self
}

func (self *ScopeForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	scope = NewScope(scope)
	ops, err := scope.Compile(self.body, nil)

	if err != nil {
		return out, err
	}

	return append(out, NewScopeOp(self, ops, scope)), nil
}

func (self *ScopeForm) Do(action func(Form) error) error {
	for _, f := range self.body {
		if err := f.Do(action); err != nil {
			return err
		}
	}

	return nil
}

func (self *ScopeForm) Dump(out io.Writer) error {
 	io.WriteString(out, "{")
	
	for i, f := range self.body {
		if i > 0 {
			io.WriteString(out, " ")
		}
		
		if err := f.Dump(out); err != nil {
			return err
		}
	}

	io.WriteString(out, "}")
	return nil
}

func (self *ScopeForm) Quote(in *Forms, scope *Scope, thread *Thread, registers *Stack, pos Pos) (Val, error) {
	out := make([]Val, len(self.body))
	var err error
	
	for i, f := range self.body {
		if out[i], err = f.Quote(nil, scope, thread, registers, pos); err != nil {
			return Nil, err
		}
	}
	
	return NewVal(&TScope, out), nil
}
