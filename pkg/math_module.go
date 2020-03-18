package gfoo

import (
	"math"
)

type MathModule struct {
	Scope
}

func divImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop().data.(Int)
	stack.Push(NewVal(&TInt, stack.Pop().data.(Int) / y))
	return nil
}

func sqrtImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, Int(math.Sqrt(float64(stack.Pop().data.(Int))))))
	return nil
}

func (self *MathModule) Init() *Scope {
	self.Scope.Init()
	self.AddMethod("div", []Arg{AType("x", &TInt), AType("y", &TInt)}, []Ret{RType(&TInt)}, divImp)
	self.AddMethod("sqrt", []Arg{AType("val", &TInt)}, []Ret{RType(&TInt)}, sqrtImp)
	return &self.Scope
}
