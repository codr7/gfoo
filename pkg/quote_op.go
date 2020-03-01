package gfoo

type QuoteOp struct {
	OpBase
}

func NewQuoteOp(form Form) *QuoteOp {
	op := new(QuoteOp)
	op.OpBase.Init(form)
	return op
}

func (self *QuoteOp) Eval(scope *Scope, stack *Slice) error {
	v, err := self.form.Quote(scope, self.form.Pos())
	
	if err != nil {
		return err
	}

	stack.Push(v)
	return nil
}
