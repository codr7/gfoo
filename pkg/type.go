package gfoo

import (
	"io"
)

type Type interface {
	Dump(val interface{}, out io.Writer) error
}

type TypeBase struct {
	id string
}

func (typ *TypeBase) Init(id string) {
	typ.id = id
}
