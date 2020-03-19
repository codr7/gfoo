package gfoo

import (
	"image/color"
)

type ImageModule struct {
	Scope
}

func rgbaNewImp(scope *Scope, stack *Slice, pos Pos) error {
	h := stack.Pop().data.(Int)
	w := stack.Pop().data.(Int)
	stack.Push(NewVal(&TRgba, NewRgba(int(w), int(h))))
	return nil
}

func rgbaSetImp(scope *Scope, stack *Slice, pos Pos) error {
	a := stack.Pop().data.(Byte)
	b := stack.Pop().data.(Byte)
	g := stack.Pop().data.(Byte)
	r := stack.Pop().data.(Byte)

	y := stack.Pop().data.(Int)
	x := stack.Pop().data.(Int)

	image := stack.Pop().data.(*Rgba)
	image.Set(int(x), int(y), color.NRGBA{R: r, G: g, B: b, A: a})
	return nil
}

func (self *ImageModule) Init() *Scope {
	self.Scope.Init()
	self.AddType(&TRgba)

	self.AddMethod("new-rgba",
		[]Arg{AType("width", &TInt), AType("height", &TInt)},
		[]Ret{RType(&TRgba)},
		rgbaNewImp)

	self.AddMethod("set",
		[]Arg{AType("image", &TRgba),
			AType("x", &TInt), AType("y", &TInt),
			AType("r", &TByte), AType("g", &TByte), AType("b", &TByte), AType("a", &TByte)},
		nil,
		rgbaSetImp)

	return &self.Scope
}
