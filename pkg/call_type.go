package gfoo

import (
	"fmt"
	"io"
)

var TCall CallType

type CallType struct {
	ValTypeBase
}

func (_ *CallType) Compare(x, y Val) Order {	
	return CompareVals(x.data.([]Val), y.data.([]Val))
}

func (_ *CallType) Dump(val Val, out io.Writer) error {
	v := val.data.([]Val)
	
	if _, err := fmt.Fprintf(out, "'%v(", v[0]); err != nil {
		return err
	}

	if err := DumpVals(v[1:], out); err != nil {
		return err
	}

	if _, err := io.WriteString(out, ")"); err != nil {
		return err
	}

	return nil
}

func (_ *CallType) New(name string, parents...Type) ValType {
	t := new(CallType)
	t.Init(name, parents...)
	return t
}

func (self *CallType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *CallType) Unquote(val Val, scope *Scope, pos Pos) Form {
	v := val.data.([]Val)
	id := TId.Unquote(v[0], scope, pos).(*Id)
	val.data = v[1:]
	return NewCallForm(id, TGroup.Unquote(val, scope, pos).(*Group), pos)
}
