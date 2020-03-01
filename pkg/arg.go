package gfoo

import (
	"io"
	"strconv"
)

type Arg struct {
	name string
	index int
	valType Type
	val Val
}

func AIndex(name string, index int) Arg {
	return Arg{name: name, index: index}
}

func AType(name string, valType Type) Arg {
	return Arg{name: name, index: -1, valType: valType}
}

func AVal(name string, val Val) Arg {
	return Arg{name: name, index: -1, val: val}
}

func (self *Arg) Dump(out io.Writer) error {
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

func (self *Arg) Match(stack []Val, index int) bool {
	if self.index == -1 {
		if self.valType == nil {
			return self.val.Compare(stack[index]) == Eq
		} else {
			xt, yt := stack[index].dataType, self.valType
			return xt == yt || xt.Isa(yt) != nil
		}
	}

	xt, yt := stack[index].dataType, stack[self.index].dataType
	return xt == yt || xt.Isa(yt) != nil
}
