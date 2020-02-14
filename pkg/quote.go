package gfoo

type Quote struct {
	FormBase
	form Form
}

func NewQuote(pos Pos, form Form) *Quote {
	return new(Quote).Init(pos, form)
}

func (self *Quote) Init(pos Pos, form Form) *Quote {
	self.FormBase.Init(pos)
	self.form = form
	return self
}

func (self *Quote) Compile(in *Forms, out []Op, vm *VM, scope *Scope) ([]Op, error) {
	v, err := self.form.Quote(vm, scope)

	if err != nil {
		return out, err
	}
	
	return append(out, NewPush(self, v)), nil
}

func (self *Quote) Quote(vm *VM, scope *Scope) (Val, error) {
	return self.form.Quote(vm, scope)
}
