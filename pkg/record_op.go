package gfoo

type RecordOp struct {
	OpBase
	fieldOps []Op
}

func NewRecordOp(form Form, fieldOps []Op) *RecordOp {
	op := new(RecordOp)
	op.OpBase.Init(form)
	op.fieldOps = fieldOps
	return op
}

func (self *RecordOp) Eval(scope *Scope, stack *Slice) error {
	var fs Slice
	
	if err := scope.EvalOps(self.fieldOps, &fs); err != nil {
		return err
	}

	r := NewRecord()

	for i := 0; i < fs.Len(); i++ {
		k := fs.items[i].data.(string)
		i++
		r.fields.Insert(k, fs.items[i], false)
	}

	stack.Push(NewVal(&TRecord, r))
	return nil
}
