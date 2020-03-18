package gfoo

import (
	"io"
	"strconv"
)

var TByte ByteType

type Byte = uint8

type ByteType struct {
	ValTypeBase
}

func (_ *ByteType) Bool(val Val) bool {
	return val.data.(Byte) != 0
}

func (_ *ByteType) Compare(x, y Val) Order {
	return CompareByte(x.data.(Byte), y.data.(Byte))
}

func (_ *ByteType) Dump(val Val, out io.Writer) error {
	_, err := io.WriteString(out, strconv.Itoa(int(val.data.(Byte))))
	return err
}

func (self *ByteType) Is(x, y Val) bool {
	return x.data == y.data
}

func (self *ByteType) Negate(val *Val) {
	val.Init(&TBool, !self.Bool(*val))
}

func (_ *ByteType) New(name string, parents...Type) ValType {
	t := new(ByteType)
	t.Init(name, parents...)
	return t
}

func (self *ByteType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ByteType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
