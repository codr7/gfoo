package gfoo

import (
	"image/png"
	"io"
)

type PngModule struct {
	Module
}

func pngEncode(scope *Scope, stack *Slice, pos Pos) error {
	image := stack.Pop().data.(*Rgba)
	out := stack.Pop().data.(io.Writer)
	return png.Encode(out, image)
}

func (self *PngModule) Init() *Module {
	self.Module.Init()

	self.AddMethod("encode", []Arg{AType("out", &TWriter), AType("image", &TRgba)}, nil, pngEncode)

	return &self.Module
}
