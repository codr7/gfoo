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

func (self *GFoo) Compile(forms []Form) ([]Op, error) {
	return nil, nil
}

func (self *GFoo) DumpStack(out io.Writer) error {
	return DumpSlice(self.stack, out)
}

func (self *GFoo) Errorf(pos Position, spec string, args...interface{}) error {
	msg := fmt.Sprintf("Error in '%v', line %v, column %v: %v ", 
		pos.filename, pos.line, pos.column, fmt.Sprintf(spec, args...))

	if self.Debug {
		panic(msg)
	}

	return errors.New(msg)
}

func (self *GFoo) Evaluate(ops []Op) error {
	return nil
}

func (self *GFoo) Parse(in *bufio.Reader, pos *Position) ([]Form, error) {
	var out []Form
	var f Form
	var err error
	
	for {
		if err = skipSpace(in, pos); err == nil {
			f, err = self.parseForm(in, pos)
		}

		if err == io.EOF {
			break
		}

		if err != nil {			
			return nil, err
		}

		out = append(out, f)
	}
	
	return out, nil
}
