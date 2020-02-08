package gfoo

type GetType struct {
	OpBase
}

func NewGetType(form Form) *GetType {
	o := new(GetType)
	o.OpBase.Init(form)
	return o
}

func (self *GetType) Evaluate(gfoo *GFoo, scope *Scope) error {
	v := gfoo.Peek()

	if v == nil {
		return gfoo.Error(self.form.Pos(), "Missing value")
	}

	v.data, v.dataType = v.dataType, &TType
	return nil
}
