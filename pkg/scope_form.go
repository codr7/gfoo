package gfoo

import (
	//"fmt"
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
	ops, err := scope.Clone().Compile(self.body, nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewScopeOp(self, ops)), nil
}

func (self *ScopeForm) Do(action func(Form) error) error {
	for _, f := range self.body {
		if err := f.Do(action); err != nil {
			return err
		}
	}

	return nil
}

func (self *ScopeForm) Quote(scope *Scope, pos Pos) (Val, error) {
	return NewVal(&TScope, self), nil
}
