package gfoo

type Drop struct {
	OpBase
}

func NewDrop(form Form) *Drop {
	o := new(Drop)
	o.OpBase.Init(form)
	return o
}

func (self *Drop) Evaluate(scope *Scope, stack *Slice) error {
	if _, ok := stack.Pop(); !ok {
		return scope.Error(self.form.Pos(), "Nothing to drop")
	}
	
	return nil
}
