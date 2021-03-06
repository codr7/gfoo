package gfoo

import (
	"fmt"
	"io"
)

type Id struct {
	FormBase
	name string
}

func NewId(name string, pos Pos) *Id {
	f := new(Id)
	f.FormBase.Init(pos)
	f.name = name
	return f
}

func (self *Id) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {	
	return self.compileName(self.name, in, out, scope, scope)
}

func (self *Id) Do(action func(Form) error) error {
	return action(self)
}

func (self *Id) Dump(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%v", self.name)
	return err
}

func (self *Id) Quote(in *Forms, scope *Scope, thread *Thread, registers []Val, pos Pos) (Val, error) {
	if in != nil {
		if next := in.Peek(); next != nil {
			g, ok := next.(*Group)

			if ok {
				in.Pop()
				return NewCallForm(self, g, pos).Quote(nil, scope, thread, registers, pos)
			}
		}
	}
	
	return NewVal(&TId, self.name), nil
}

func compileArgs(in *Forms, scope *Scope) ([]Op, error) {
	if next := in.Peek(); next != nil {
		if _, ok := next.(*Group); ok {
			return in.Pop().Compile(nil, nil, scope)
		}
	}

	return nil, nil
}

func (self *Id) compileFunction(f *Function, in *Forms, out []Op, scope *Scope) ([]Op, error) {	
	if len(f.methods) == 1 {
		return self.compileMethod(f.methods[0], in, out, scope)
	}

	var args []Op
	var err error
	
	if args, err = compileArgs(in, scope); err != nil {
		return out, err
	}
	
	v := NewVal(&TFunction, f)
	return append(out, NewCall(self, &v, args)), nil
}

func (self *Id) compileMethod(m *Method, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	var args []Op
	var err error
	
	if args, err = compileArgs(in, scope); err != nil {
		return out, err
	}

	v := NewVal(&TMethod, m)
	return append(out, NewCall(self, &v, args)), nil
}

func (self *Id) compileName(
	name string,
	in *Forms,
	out []Op,
	nameScope, scope *Scope) ([]Op, error) {		
	b := nameScope.Get(name)
	
	if  b == nil {
		if name[0] == '.' {
			return append(out, NewGet(self, name[1:])), nil
		}
		
		return out, Error(self.pos, "Unknown identifier: %v", name)
	}
	
	if b.val == Undefined {
		return append(out, NewLoad(self, name, nameScope.registers[name])), nil
	}
	
	v := &b.val

	switch (v.dataType) {
	case &TFunction:
		return self.compileFunction(v.data.(*Function), in, out, scope)
	case &TMacro:
		return v.data.(*Macro).Expand(self, in, out, scope)
	case &TMethod:
		return self.compileMethod(v.data.(*Method), in, out, scope)
	case &TModule:
		if next := in.Peek(); next != nil {
			if id, ok := next.(*Id); ok && id.name[0] == '.' {
				in.Pop()
				
				return self.compileName(
					id.name[1:],
					in,
					out,
					&v.data.(*Module).Scope,
					scope)
			}
		}
	}
	
	return append(out, NewPush(self, *v)), nil
}
