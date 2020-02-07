package gfoo

import (
	"io"
)

type Type interface {
	Compare(x, y interface{}) Order
	Dump(val interface{}, out io.Writer) error
	Name() string
	Unquote(val interface{}) Form
}

type TypeBase struct {
	name string
}

func (self *TypeBase) Init(name string) {
	self.name = name
}

func (self *TypeBase) Name() string {
	return self.name
}
