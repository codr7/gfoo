package gfoo

type Drop struct {
	OpBase
}

func NewDrop(form Form) *Drop {
	o := new(Drop)
	o.OpBase.Init(form)
	return o
}

func (self *Drop) Evaluate(stack *Slice, scope *Scope) error {
	if _, ok := stack.Pop(); !ok {
		return scope.vm.Error(self.form.Pos(), "Nothing to drop")
	}
	
	return nil
}
