package gfoo

import (
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
	if b := scope.Get(self.name); b != nil && (self.name == "NIL" || b.val != Nil) {
		v := &b.val
		
		switch (v.dataType) {
		case &TFunction:
			return self.compileFunction(v.data.(*Function), out)
		case &TMacro:
			return v.data.(*Macro).Expand(self, in, out, scope)
		case &TMethod:
			return self.compileMethod(v.data.(*Method), out)
		}
		
		return append(out, NewPush(self, *v)), nil
	}

	return append(out, NewGet(self, self.name)), nil
}

func (self *Id) Do(action func(Form) error) error {
	return action(self)
}

func (self *Id) Dump(out io.Writer) error {
	_, err := io.WriteString(out, self.name)
	return err
}

func (self *Id) Quote(scope *Scope, pos Pos) (Val, error) {
	return NewVal(&TId, self.name), nil
}

func (self *Id) compileFunction(f *Function, out []Op) ([]Op, error) {	
	if len(f.methods) == 1 {
		return self.compileMethod(f.methods[0], out)
	}
	
	v := NewVal(&TFunction, f)
	return append(out, NewCall(self, &v, nil)), nil
}

func (self *Id) compileMethod(m *Method, out []Op) ([]Op, error) {
	v := NewVal(&TMethod, m)
	return append(out, NewCall(self, &v, nil)), nil
}
