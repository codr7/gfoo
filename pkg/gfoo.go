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
	stack Slice
}

func dropImp(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	return append(out, NewDrop(form)), nil
}

func dupImp(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	return append(out, NewDup(form)), nil
}

func resetImp(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	return append(out, NewReset(form)), nil
}

func letImp(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	key, ok := args.Pop().(*Id)

	if !ok {
		gfoo.Error(key.Pos(), "Expected id: %v", key)
	}

	if found := scope.Get(key.name); found == nil {
		scope.Set(key.name, nil, nil)
	} else {
		if found.scope == scope {
			return out, gfoo.Error(key.Pos(), "Duplicate binding: %v", key.name) 
		}
	}
	
	val := args.Pop()
	
	if id, ok := val.(*Id); !ok || id.name != "_" {
		var err error
		
		if out, err = val.Compile(gfoo, scope, &NilForms, out); err != nil {
			return out, err
		}
	}
	
	return append(out, NewLet(form, key.name)), nil
}

func typeImp(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	return append(out, NewGetType(form)), nil
}
	
func New() *GFoo {
	g := new(GFoo)
	g.rootScope.Init()
	
	g.AddConst("T", &TBool, true)
	g.AddConst("F", &TBool, false)

	g.AddMacro("_", 0, dropImp)
	g.AddMacro("..", 0, dupImp)
	g.AddMacro("|", 0, resetImp)
	g.AddMacro("let:", 0, letImp)
	g.AddMacro("type", 0, typeImp)
	return g
}

func (self *GFoo) AddConst(name string, dataType Type, data interface{}) {
	self.rootScope.Set(name, dataType, data)
}

func (self *GFoo) AddMacro(name string, argCount int, imp MacroImp) {
	self.rootScope.Set(name, &TMacro, NewMacro(name, argCount, imp))
}

func (self *GFoo) Compile(in []Form, scope *Scope, out []Op) ([]Op, error) {
	var err error
	var inForms Forms
	inForms.Init(in)
	
	for f := inForms.Pop(); f != nil; f = inForms.Pop() {
		if out, err = f.Compile(self, scope, &inForms, out); err != nil {
			return out, err
		}
	}
	
	return out, nil
}

func (self *GFoo) StackDump(out io.Writer) error {
	return self.stack.Dump(out)
}

func (self *GFoo) Error(pos Pos, spec string, args...interface{}) error {
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

func (self *GFoo) Peek() *Val {
	return self.stack.Peek()
}

func (self *GFoo) Pop() *Val {
	return self.stack.Pop()
}

func (self *GFoo) Push(dataType Type, data interface{}) {
	self.stack.Push(dataType, data)
}

func (self *GFoo) RootScope() *Scope {
	return &self.rootScope
}
