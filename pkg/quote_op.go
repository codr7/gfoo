package gfoo

type QuoteOp struct {
	OpBase
	scope *Scope
}

func NewQuoteOp(form Form, scope *Scope) *QuoteOp {
	op := new(QuoteOp)
	op.OpBase.Init(form)
	op.scope = scope
	return op
}

func (self *QuoteOp) Eval(thread *Thread, registers, stack *Stack) error {
	v, err := self.form.Quote(nil, self.scope, thread, registers, self.form.Pos())
	
	if err != nil {
		return err
	}

	stack.Push(v)
	return nil
}
