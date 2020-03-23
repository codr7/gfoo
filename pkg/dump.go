package gfoo

import (
	"fmt"
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

func DumpVals(in []Val, out io.Writer) error {
	for i, v := range in {
		if i > 0 {
			if _, err := fmt.Fprint(out, " "); err != nil {
				return err
			}
		}
		
		if err := v.Dump(out); err != nil {
			return err
		}
	}
	
	return nil
}
