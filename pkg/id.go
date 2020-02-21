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
	if b := scope.Get(self.name); b != nil && b.val != Nil {
		v := &b.val

		switch (v.dataType) {
		case &TMacro:
			return v.data.(*Macro).Expand(self, in, out, scope)
		case &TMethod:
			v := NewVal(&TMethod, v.data.(*Method));
			return append(out, NewCall(self, &v, nil)), nil
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
