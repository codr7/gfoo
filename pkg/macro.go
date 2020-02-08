package gfoo

import (
	"fmt"
	"io"
)

var TMacro MacroType

func init() {
	TMacro.Init("Macro")
}

type MacroType struct {
	TypeBase
}

func (_ *MacroType) Compare(x, y interface{}) Order {
	return CompareString(x.(*Macro).name, y.(*Macro).name)
}

func (_ *MacroType) Dump(val interface{}, out io.Writer) error {
	_, err := fmt.Fprintf(out, "Macro(%v)", val.(*Macro).name)
	return err
}

func (self *MacroType) Unquote(pos Pos, val interface{}) Form {
	return NewLiteral(pos, self, val)
}

type MacroImp = func(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error)

type Macro struct {
	name string
	argCount int
	imp MacroImp
}

func NewMacro(name string, argCount int, imp MacroImp) *Macro {
	return &Macro{name: name, argCount: argCount, imp: imp}
}

func (self *Macro) Expand(gfoo *GFoo, scope *Scope, form Form, args *Forms, out []Op) ([]Op, error) {
	if args.Len() < self.argCount {
		gfoo.Error(form.Pos(), "Not enough arguments: %v", self.name)
	}
	
	return self.imp(gfoo, scope, form, args, out)
}

func (self *Scope) SetMacro(name string, argCount int, imp MacroImp) {
	self.Set(name, &TMacro, NewMacro(name, argCount, imp))
}
