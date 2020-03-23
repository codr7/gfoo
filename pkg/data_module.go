package gfoo

type DataModule struct {
	Module
}

func recordImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var fields *Group
	var ok bool

	if fields, ok = f.(*Group); !ok {
		return out, Error(form.Pos(), "Invalid fields: %v", f)
	}

	fieldForms := NewForms(fields.body)
	var fieldOps []Op

	for {
		if f = fieldForms.Pop(); f == nil {
			break
		}

		id, ok := f.(*Id)
		
		if !ok {
			return out, Error(f.Pos(), "Expected id: %v", f)
		}

		fieldOps = append(fieldOps, NewPush(f, NewVal(&TId, id.name)))
		var err error

		if f = fieldForms.Pop(); f == nil {
			return out, Error(id.Pos(), "Missing value: %v", id)
		}
		
		if fieldOps, err = f.Compile(fieldForms, fieldOps, scope); err != nil {
			return out, err
		}
	}

	return append(out, NewRecordOp(form, fieldOps)), nil
}

func recordFieldsImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	in, err := stack.Pop().Iter(pos)

	if err != nil {
		return err
	}

	stack.Push(NewVal(&TIter, in))
	return nil
}

func recordLengthImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, Int(stack.Pop().data.(*Record).Len())))
	return nil
}

func recordMergeImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	source := stack.Pop().data.(*Record)
	stack.Pop().data.(*Record).Merge(source)
	return nil
}

func recordSetImp(thread *Thread, registers, stack *Slice, pos Pos) error {
	v, k, r := stack.Pop(), stack.Pop(), stack.Pop()
	r.data.(*Record).Set(k.data.(string), *v)
	return nil
}

func (self *DataModule) Init() *Module {
	self.Module.Init()

	self.AddType(&TRecord)
	self.AddMacro("record:", 1, recordImp)

	self.AddMethod("fields", []Arg{AType("val", &TRecord)}, []Ret{RType(&TIter)}, recordFieldsImp)
	self.AddMethod("length", []Arg{AType("val", &TRecord)}, []Ret{RType(&TInt)}, recordLengthImp)
	self.AddMethod("merge", []Arg{AType("target", &TRecord), AType("source", &TRecord)}, nil, recordMergeImp)

	self.AddMethod("set",
		[]Arg{AType("record", &TRecord), AType("key", &TId), AType("val", &TAny)},
		nil,
		recordSetImp)

	return &self.Module
}
