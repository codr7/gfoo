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
	
	rootScope Scope
	stack []Val
}

func New() *GFoo {
	g := new(GFoo)
	g.rootScope.Init()
	return g
}

func (self *GFoo) Compile(in []Form, scope *Scope, out []Op) ([]Op, error) {
	var err error
	
	for _, f := range in {
		if out, err = f.Compile(self, scope, out); err != nil {
			return out, err
		}
	}
	
	return out, nil
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

func (self *GFoo) Evaluate(ops []Op, scope *Scope) error {
	for _, o := range ops {
		if err := o.Evaluate(self, scope); err != nil {
			return err
		}
	}
	
	return nil
}

func (self *GFoo) Parse(in *bufio.Reader, pos *Position, out []Form) ([]Form, error) {
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
			return out, err
		}

		out = append(out, f)
	}
	
	return out, nil
}

func (self *GFoo) Push(val Val) {
	self.stack = append(self.stack, val)
}

func (self *GFoo) RootScope() *Scope {
	return &self.rootScope
}
