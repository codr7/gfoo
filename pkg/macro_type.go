package gfoo

import (
	"fmt"
	"io"
)

var TMacro MacroType

func init() {
	TMacro.Init("Macro")
}

type MacroType struct {
	TypeBase
}

func (_ *MacroType) Compare(x, y interface{}) Order {
	return CompareString(x.(*Macro).name, y.(*Macro).name)
}

func (_ *MacroType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Macro(%v)", val.(*Macro).name)
	return err
}

func (self *MacroType) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, self, val)
}
