package gfoo

import (
	"fmt"
	"io"
)

var TMacro MacroType

type MacroType struct {
	ValTypeBase
}

func (_ *MacroType) Compare(x, y Val) Order {
	return CompareString(x.data.(*Macro).name, y.data.(*Macro).name)
}

func (_ *MacroType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Macro(%v)", val.data.(*Macro).name)
	return err
}

func (self *MacroType) Print(val Val, out io.Writer) error {
	return self.Dump(val, out)
}

func (self *MacroType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val, pos)
}
