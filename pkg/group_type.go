package gfoo

import (
	"io"
)

var TGroup GroupType

type GroupType struct {
	ValTypeBase
}

func (_ *GroupType) Compare(x, y Val) Order {	
	return CompareVals(x.data.([]Val), y.data.([]Val))
}

func (_ *GroupType) Dump(val Val, out io.Writer) error {
	if _, err := io.WriteString(out, "'("); err != nil {
		return err
	}

	if err := DumpVals(val.data.([]Val), out); err != nil {
		return err
	}

	if _, err := io.WriteString(out, ")"); err != nil {
		return err
	}

	return nil
}

func (_ *GroupType) New(name string, parents...Type) ValType {
	t := new(GroupType)
	t.Init(name, parents...)
	return t
}

func (self *GroupType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *GroupType) Unquote(val Val, scope *Scope, pos Pos) Form {
	in := val.data.([]Val)
	out := make([]Form, len(in))
	
	for i, v := range in {
		out[i] = v.Unquote(scope, pos);
	}

	return NewGroup(out, pos)
}
