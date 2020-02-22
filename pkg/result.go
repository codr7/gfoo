package gfoo

import (
	"io"
	"strconv"
)

type Result struct {
	index int
	valType Type
	val Val
}

func RIndex(index int) Result {
	return Result{index: index}
}

func RType(valType Type) Result {
	return Result{index: -1, valType: valType}
}

func RVal(val Val) Result {
	return Result{index: -1, val: val}
}

func (self Result) Dump(out io.Writer) error {
	if self.index == -1 {
		if self.valType == nil {
			if err := self.val.Dump(out); err != nil {
				return err
			}
		} else {
			if _, err := io.WriteString(out, self.valType.Name()); err != nil {
				return err
			}
		}
	} else {
		if _, err := io.WriteString(out, strconv.Itoa(self.index)); err != nil {
			return err
		}
	}

	return nil
}
