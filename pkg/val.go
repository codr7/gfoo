package gfoo

import (
	"io"
	"strings"
)

type Val struct {
	dataType Type
	data interface{}
}

func NewVal(dataType Type, data interface{}) Val {
	var v Val
	v.Init(dataType, data)
	return v
}

func (self *Val) Init(dataType Type, data interface{}) {
	self.dataType = dataType
	self.data = data
}

func (self Val) Compare(other Val) Order {
	if self.dataType != other.dataType {
		return strings.Compare(self.dataType.Name(), other.dataType.Name())
	}

	return self.dataType.Compare(self.data, other.data)
}

func (self Val) Dump(out io.Writer) error {
	return self.dataType.Dump(self.data, out)
}

func (self Val) Is(other Val) bool {
	if self.dataType != other.dataType {
		return false
	}
	
	return self.data == other.data
}

func (self Val) Literal(pos Pos) *Literal {
	return NewLiteral(pos, self.dataType, self.data)
}

func (self Val) Unquote(pos Pos) Form {
	return self.dataType.Unquote(pos, self.data)
}
