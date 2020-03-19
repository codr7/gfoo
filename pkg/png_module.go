package gfoo

import (
	"image/png"
	"io"
)

type PngModule struct {
	Module
}

func pngEncode(scope *Scope, stack *Slice, pos Pos) error {
	out := stack.Pop().data.(io.Writer)
	image := stack.Pop().data.(*Rgba)
	return png.Encode(out, image)
}

func (self *PngModule) Init() *Module {
	self.Module.Init()

	self.AddMethod("encode", []Arg{AType("image", &TRgba), AType("out", &TWriter)}, nil, pngEncode)

	return &self.Module
}
