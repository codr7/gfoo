package gfoo

import (
	"io"
)

type Value struct {
	dataType Type
	data interface{}
}

func NewValue(typ Type, dat interface{}) *Value {
	return &Value{dataType: typ, data: dat} 
}

func (val *Value) Dump(out io.Writer) error {
	return val.dataType.Dump(val.data, out)
}
