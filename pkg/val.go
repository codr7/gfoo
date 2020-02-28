package gfoo

import (
	"io"
	"strings"
)

var Nil Val

type Val struct {
	dataType ValType
	data interface{}
}

func NewVal(dataType ValType, data interface{}) Val {
	var v Val
	v.Init(dataType, data)
	return v
}

func (self *Val) Init(dataType ValType, data interface{}) {
	self.dataType = dataType
	self.data = data
}

func (self Val) Bool() bool {
	return self.dataType.Bool(self)
}

func (self Val) Call(scope *Scope, stack *Slice, pos Pos) error {
	return self.dataType.Call(self, scope, stack, pos)
}

func (self Val) Clone() Val {
	return NewVal(self.dataType, self.dataType.Clone(self))
}

func (self Val) Compare(other Val) Order {
	if self.dataType != other.dataType {
		return strings.Compare(self.dataType.Name(), other.dataType.Name())
	}

	return self.dataType.Compare(self, other)
}

func (self Val) Dump(out io.Writer) error {
	return self.dataType.Dump(self, out)
}

func (self Val) Get(key string, scope *Scope, pos Pos) (Val, error) {
	return self.dataType.Get(self, key, scope, pos)
}

func (self Val) Is(other Val) bool {
	if self.dataType != other.dataType {
		return false
	}
	
	return self.data == other.data
}

func (self Val) Literal(pos Pos) *Literal {
	return NewLiteral(self, pos)
}

func (self Val) Print(out io.Writer) error {
	return self.dataType.Print(self, out)
}

func (self Val) Unquote(scope *Scope, pos Pos) Form {
	return self.dataType.Unquote(self, scope, pos)
}
