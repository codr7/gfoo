package gfoo

type Quote struct {
	FormBase
	form Form
}

func NewQuote(form Form, pos Pos) *Quote {
	return new(Quote).Init(form, pos)
}

func (self *Quote) Init(form Form, pos Pos) *Quote {
	self.FormBase.Init(pos)
	self.form = form
	return self
}

func (self *Quote) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	v, err := self.form.Quote(scope)

	if err != nil {
		return out, err
	}
	
	return append(out, NewPush(self, v)), nil
}

func (self *Quote) Do(action func(Form) error) error {
	return self.form.Do(action)
}

func (self *Quote) Quote(scope *Scope) (Val, error) {
	return self.form.Quote(scope)
}
