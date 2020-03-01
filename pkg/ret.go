package gfoo

import (
	"io"
	"strconv"
)

type Ret struct {
	index int
	valType Type
	val Val
}

func RIndex(index int) Ret {
	return Ret{index: index}
}

func RType(valType Type) Ret {
	return Ret{index: -1, valType: valType}
}

func RVal(val Val) Ret {
	return Ret{index: -1, val: val}
}

func (self Ret) Dump(out io.Writer) error {
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
