package gfoo

import (
	"fmt"
	"io"
)

var TId SymbolType

func init() {
	TId.Init("Id")
}

type SymbolType struct {
	StringType
}

func (_ *SymbolType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "'%v", val.data.(string))
	return err
}

func (_ *SymbolType) Unquote(val Val, pos Pos) Form {
	return NewId(val.data.(string), pos)
}
