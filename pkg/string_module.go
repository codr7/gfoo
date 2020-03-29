package gfoo

import (
	"strings"
)

type StringModule struct {
	Module
}

func charChainImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	y := stack.Pop().data.(rune)
	x := stack.Pop().data.(rune)
	stack.Push(NewVal(&TString, string([]rune{x, y})))
	return nil
}

func stringChainImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	y := stack.Pop().data.(string)
	x := stack.Pop().data.(string)
	stack.Push(NewVal(&TString, string(append([]rune(x), []rune(y)...))))
	return nil
}

func stringDownImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	stack.Push(NewVal(&TString, strings.ToLower(stack.Pop().data.(string))))
	return nil
}

func stringUpImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	stack.Push(NewVal(&TString, strings.ToUpper(stack.Pop().data.(string))))
	return nil
}

func (self *StringModule) Init() *Module {
	self.Module.Init()
	
	self.AddMethod("~", []Arg{AType("x", &TChar), AType("y", &TChar)}, []Ret{RType(&TString)}, charChainImp)

	self.AddMethod("~",
		[]Arg{AType("x", &TString), AType("y", &TString)},
		[]Ret{RType(&TString)},
		stringChainImp)

	self.AddMethod("down", []Arg{AType("in", &TString)}, []Ret{RType(&TString)}, stringDownImp)
	self.AddMethod("up", []Arg{AType("in", &TString)}, []Ret{RType(&TString)}, stringUpImp)

	return &self.Module
}
