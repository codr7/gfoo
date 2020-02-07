package gfoo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	return DumpSlice(gfoo.stack, out)
}

func (gfoo *GFoo) Errorf(pos Position, spec string, args...interface{}) error {
	msg := fmt.Sprintf("Error in '%v', line %v, column %v: %v ", 
		pos.filename, pos.line, pos.column, fmt.Sprintf(spec, args...))

	if gfoo.Debug {
		panic(msg)
	}

	return errors.New(msg)
}

func (gfoo *GFoo) Evaluate(ops []Op) error {
	return nil
}

func (gfoo *GFoo) Parse(in *bufio.Reader, pos *Position) ([]Form, error) {
	var out []Form
	var f Form
	var err error
	
	for {
		if f, err = gfoo.parseForm(in, pos); err != nil {
			if err == io.EOF {
				break
			}
			
			return nil, err
		}

		out = append(out, f)
	}
	
	return out, nil
}
