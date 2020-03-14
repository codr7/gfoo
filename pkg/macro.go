package gfoo

type MacroImp = func(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error)

type Macro struct {
	name string
	argCount int
	imp MacroImp
}

func NewMacro(name string, argCount int, imp MacroImp) *Macro {
	return &Macro{name: name, argCount: argCount, imp: imp}
}

func (self *Macro) Expand(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	inLen := 0

	if in != nil {
		inLen = in.Len()
	}
	
	if inLen < self.argCount {
		scope.Error(form.Pos(), "Not enough arguments: %v (%v)", inLen, self.argCount)
	}
	
	return self.imp(form, in, out, scope)
}
