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
	
	symbols sync.Map
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

func (gfoo *GFoo) Eval(ops []Op) error {
	return nil
}

func (gfoo *GFoo) Parse(source string) ([]Form, error) {
	return nil, nil
}

func (gfoo *GFoo) Symbol(name string) Value {
	found, _ := gfoo.symbols.Load(name)

	if found != nil {
		return found.(Value)
	}
	
	s := NewValue(&Symbol, name)
	gfoo.symbols.Store(name, s)
	return s
}
