package gfoo

import (
	"fmt"
)

type UnionType struct {
	Trait
	types []Type
}

func Union(name string, types...Type) *UnionType {
	return new(UnionType).Init(name, types...)
}

func (self *UnionType) Init(name string, types...Type) *UnionType {
	self.Trait.Init(name)
	self.types = types
	return self
}

func Option(in Type) Type {
	return Union(fmt.Sprintf("%v?", in.Name()), in, &TNil)
}
