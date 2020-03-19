package gfoo

import (
	"strings"
)

type StringModule struct {
	Module
}

func stringDownImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TString, strings.ToLower(stack.Pop().data.(string))))
	return nil
}

func stringUpImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TString, strings.ToUpper(stack.Pop().data.(string))))
	return nil
}

func (self *StringModule) Init() *Module {
	self.Module.Init()
	
	self.AddMethod("down", []Arg{AType("in", &TString)}, []Ret{RType(&TString)}, stringDownImp)
	self.AddMethod("up", []Arg{AType("in", &TString)}, []Ret{RType(&TString)}, stringUpImp)

	return &self.Module
}
