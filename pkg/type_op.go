package gfoo

type TypeOp struct {
	OpBase
}

func NewTypeOp(form Form) *TypeOp {
	op := new(TypeOp)
	op.OpBase.Init(form)
	return op
}

func (self *TypeOp) Evaluate(scope *Scope, stack *Slice) error {
	v := stack.Peek()

	if v == nil {
		return scope.Error(self.form.Pos(), "Missing value")
	}

	v.data, v.dataType = v.dataType, &TMeta
	return nil
}
