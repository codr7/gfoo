package gfoo

type RecordOp struct {
	OpBase
	fields []Op
}

func NewRecordOp(form Form, fields []Op) *RecordOp {
	op := new(RecordOp)
	op.OpBase.Init(form)
	op.fields = fields
	return op
}

func (self *RecordOp) Eval(thread *Thread, registers []Val, stack *Stack) error {
	var fs Stack
	
	if err := EvalOps(self.fields, thread, registers, &fs); err != nil {
		return err
	}

	r := NewRecord()

	for i := 0; i < fs.Len(); i++ {
		k := fs.items[i].data.(string)
		i++
		r.Set(k, fs.items[i])
	}

	stack.Push(NewVal(&TRecord, r))
	return nil
}
