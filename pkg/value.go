package gfoo

import (
	"io"
	"strings"
)

type Value struct {
	dataType Type
	data interface{}
}

func NewValue(dataType Type, data interface{}) Value {
	var v Value
	v.Init(dataType, data)
	return v
}

func (self *Value) Init(dataType Type, data interface{}) {
	self.dataType = dataType
	self.data = data
}

func (self Value) Compare(other Value) Order {
	if self.dataType != other.dataType {
		return strings.Compare(self.dataType.Name(), other.dataType.Name())
	}

	return self.dataType.Compare(self.data, other.data)
}

func (self Value) Dump(out io.Writer) error {
	return self.dataType.Dump(self.data, out)
}

func (self Value) Is(other Value) bool {
	if self.dataType != other.dataType {
		return false
	}
	
	return self.data == other.data
}

func (self Value) Literal() *Literal {
	return &Literal{value: self}
}

func (self Value) Unquote() Form {
	return self.dataType.Unquote(self.data)
}
