package gfoo

import (
	"strings"
)

func stringDownImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TString, strings.ToLower(stack.Pop().data.(string))))
	return nil
}

func stringUpImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TString, strings.ToUpper(stack.Pop().data.(string))))
	return nil
}

func (self *Scope) InitStringModule() *Scope {
	self.AddMethod("down", []Arg{AType("in", &TString)}, []Ret{RType(&TString)}, stringDownImp)
	self.AddMethod("up", []Arg{AType("in", &TString)}, []Ret{RType(&TString)}, stringUpImp)
	return self
}
