package gfoo

import (
	"image"
)

type Rgba = image.NRGBA

func NewRgba(w, h int) *Rgba {
	return image.NewNRGBA(image.Rect(0, 0, w, h))
}
