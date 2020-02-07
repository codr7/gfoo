package gfoo

import (
	"io"
)

type Type interface {
	Compare(x, y interface{}) Order
	Dump(val interface{}, out io.Writer) error
	Name() string
}

type TypeBase struct {
	name string
}

func (typ *TypeBase) Init(name string) {
	typ.name = name
}

func (typ *TypeBase) Name() string {
	return typ.name
}
