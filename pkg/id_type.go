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

func (_ *SymbolType) Unquote(val Val, scope *Scope, pos Pos) Form {
	n := val.data.(string)
	
	if n[0] == '$' {
		n = scope.Unique(n)
	}

	return NewId(n, pos)
}
