package gfoo

import (
	"fmt"
	"io"
)

var TScope ScopeType

type ScopeType struct {
	ValTypeBase
}

func (_ *ScopeType) Compare(x, y Val) Order {	
	return CompareVals(x.data.([]Val), y.data.([]Val))
}

func (_ *ScopeType) Dump(val Val, out io.Writer) error {
	if _, err := io.WriteString(out, "'{"); err != nil {
		return err
	}

	if err := DumpVals(val.data.([]Val), out); err != nil {
		return err
	}

	if _, err := io.WriteString(out, "}"); err != nil {
		return err
	}

	return nil
}

func (_ *ScopeType) New(name string, parents...Type) ValType {
	t := new(ScopeType)
	t.Init(name, parents...)
	return t
}

func (self *ScopeType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *ScopeType) Unquote(val Val, scope *Scope, pos Pos) Form {
	in := val.data.([]Val)
	out := make([]Form, len(in))
	
	for i, v := range in {
		out[i] = v.Unquote(scope, pos)
	}

	fmt.Printf("%v\n", DumpString(NewScopeForm(out, pos)))
	return NewScopeForm(out, pos)
}
