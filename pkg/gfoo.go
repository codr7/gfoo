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
	
	g.rootScope.Set("T", &Bool, true)
	g.rootScope.Set("F", &Bool, false)
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

func (self *GFoo) Errorf(pos Pos, spec string, args...interface{}) error {
	msg := fmt.Sprintf("Error in '%v', line %v, column %v: %v ", 
		pos.source, pos.line, pos.column, fmt.Sprintf(spec, args...))

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

func (self *GFoo) Let(scope *Scope, pos Pos, key string, dataType Type, data interface{}) {
	if found := scope.Get(key); found == nil {
		if found.scope == scope {
			self.Errorf(pos, "Duplicate binding: %v", key) 
		}
	} else {
		scope.Set(key, dataType, data)
	}
}

func (self *GFoo) Parse(in *bufio.Reader, pos *Pos, out []Form) ([]Form, error) {
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
