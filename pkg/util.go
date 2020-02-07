package gfoo

import (
	"fmt"
	"io"
)

func MinInt(x, y int) int {
	if y < x {
		return y
	}

	return x
}

func DumpSlice(in []Value, out io.Writer) error {
	if _, err := fmt.Fprint(out, "["); err != nil {
		return err
	}

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
	
	if _, err := fmt.Fprint(out, "]"); err != nil {
		return err
	}
	
	return nil
}

