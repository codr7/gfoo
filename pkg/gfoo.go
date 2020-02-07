package gfoo

import (
	"fmt"
	"io"
	"sync"
)

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 1
)
	
type GFoo struct {
	Debug bool

	stack []Value
}

func New() *GFoo {
	gfoo := new(GFoo)
	return gfoo
}

func (gfoo *GFoo) Compile(forms []Form) ([]Op, error) {
	return nil, nil
}

func (gfoo *GFoo) DumpStack(out io.Writer) error {
	if _, err := fmt.Fprint(out, "["); err != nil {
		return err
	}

	for i, v := range gfoo.stack {
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

func (gfoo *GFoo) Evaluate(ops []Op) error {
	return nil
}

func (gfoo *GFoo) Parse(source string, pos *Position) ([]Form, error) {
	return nil, nil
}
