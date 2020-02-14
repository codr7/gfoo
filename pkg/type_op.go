package gfoo

type TypeOp struct {
	OpBase
}

func NewTypeOp(form Form) *TypeOp {
	o := new(TypeOp)
	o.OpBase.Init(form)
	return o
}

func (self *TypeOp) Evaluate(stack *Slice, scope *Scope) error {
	v := stack.Peek()

	if v == nil {
		return scope.vm.Error(self.form.Pos(), "Missing value")
	}

	v.data, v.dataType = v.dataType, &TMeta
	return nil
}
