package gfoo

import (
	"io"
	"strings"
)

var Nil, Undefined Val

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

func (self Val) Call(thread *Thread, stack *Stack, pos Pos) error {
	tt, ok := self.dataType.(TargetType)

	if !ok {
		return Error(pos, "Calling is only supported for target types: %v", self.dataType.Name())
	}

	return tt.Call(self, thread, stack, pos)
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

func (self Val) Get(key string, pos Pos) (Val, error) {
	return self.dataType.Get(self, key, pos)
}

func (self Val) Is(other Val) bool {
	if self.dataType != other.dataType {
		return false
	}
	
	return self.dataType.Is(self, other)
}

func (self Val) Iter(pos Pos) (Iter, error) {
	st, ok := self.dataType.(SequenceType)

	if !ok {
		return nil, Error(pos, "Iteration is only supported for sequence types: %v", self.dataType.Name())
	}
	
	return st.Iter(self, pos)
}

func (self Val) Keys() []string {
	return self.dataType.Keys(self)
}

func (self Val) Literal(pos Pos) *Literal {
	return NewLiteral(self, pos)
}

func (self *Val) Negate() {
	self.dataType.Negate(self)
}


func (self Val) Print(out io.Writer) error {
	return self.dataType.Print(self, out)
}

func (self Val) Unquote(scope *Scope, pos Pos) Form {
	return self.dataType.Unquote(self, scope, pos)
}
