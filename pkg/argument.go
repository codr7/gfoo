package gfoo

import (
	"fmt"
	"io"
)

type Argument struct {
	id string
	val Val
}

func NewArgument(id string, val Val) Argument {
	var a Argument
	return a.Init(id, val)
}

func (self Argument) Init(id string, val Val) Argument {
	self.id = id
	self.val = val
	return self
}

func (self Argument) Dump(out io.Writer) error {
	if self.id != "" {
		if _, err := fmt.Fprintf(out, "%v ", self.id); err != nil {
			return err
		}
	}

	if self.val.data == nil {
		if t := self.val.dataType; t != nil {
			if _, err := io.WriteString(out, t.Name()); err != nil {
				return err
			}
		}
	} else {
		if err := self.val.Dump(out); err != nil {
			return err
		}
	}

	return nil
}
