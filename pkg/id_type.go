package gfoo

import (
	"fmt"
	"io"
)

var TId IdType

type IdType struct {
	StringType
}

func (_ *IdType) Dump(val Val, out io.Writer) error {
	_, err := fmt.Fprintf(out, "'%v", val.data.(string))
	return err
}

func (_ *IdType) Unquote(val Val, scope *Scope, pos Pos) Form {
	n := val.data.(string)
	
	if n[0] == '$' {
		n = scope.Unique(n)
	}
	
	return NewId(n, pos)
}
