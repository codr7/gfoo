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
	var val Value
	val.Init(dataType, data)
	return val
}

func (val *Value) Init(dataType Type, data interface{}) {
	val.dataType = dataType
	val.data = data
}

func (val Value) Compare(other Value) Order {
	if val.dataType != other.dataType {
		return strings.Compare(val.dataType.Name(), other.dataType.Name())
	}

	return val.dataType.Compare(val.data, other.data)
}

func (val Value) Dump(out io.Writer) error {
	return val.dataType.Dump(val.data, out)
}

func (val Value) Is(other Value) bool {
	if val.dataType != other.dataType {
		return false
	}
	
	return val.data == other.data
}
