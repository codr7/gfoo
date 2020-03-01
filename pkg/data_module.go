package gfoo

func recordImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var fields *Group
	var ok bool

	if fields, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid field list: %v", f)
	}

	fieldForms := NewForms(fields.body)
	var fieldOps []Op

	for {
		f = fieldForms.Pop()

		if f == nil {
			break
		}

		id, ok := f.(*Id)
		
		if !ok {
			return out, scope.Error(f.Pos(), "Invalid field id: %v", f)
		}

		fieldOps = append(fieldOps, NewPush(f, NewVal(&TId, id.name)))
		var err error

		f = fieldForms.Pop()

		if f == nil {
			return out, scope.Error(id.Pos(), "Missing field value: %v", id)
		}
		
		if fieldOps, err = f.Compile(fieldForms, fieldOps, scope); err != nil {
			return out, err
		}
	}

	return append(out, NewRecordOp(form, fieldOps)), nil
}

func recordLengthImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TInt, NewInt(int64(stack.Pop().data.(Record).Len()))))
	return nil
}

func recordSetImp(stack *Slice, scope *Scope, pos Pos) error {
	v, k, r := stack.Pop(), stack.Pop(), stack.Pop()
	stack.Push(NewVal(&TRecord, r.data.(Record).Set(k.data.(string), *v)))
	return nil
}

func (self *Scope) InitData() *Scope {
	self.AddType(&TRecord)

	self.AddMacro("record:", 1, recordImp)

	self.AddMethod("length", []Arg{AType("val", &TRecord)}, []Ret{RType(&TInt)}, recordLengthImp)

	self.AddMethod("set",
		[]Arg{AType("record", &TRecord), AType("key", &TId), AType("val", &TAny)},
		[]Ret{RType(&TRecord)},
		recordSetImp)

	return self
}
