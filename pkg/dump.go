package gfoo

import (
	"io"
	"strings"
)

type Dumper interface {
	Dump(out io.Writer) error
}

func DumpString(in Dumper) string {
	var out strings.Builder
	in.Dump(&out)
	return out.String()
}
