package gfoo

func recordImp(form Form, in *Forms, out []Op, scope *Scope) ([]Op, error) {
	f := in.Pop()
	var fields *Group
	var ok bool

	if fields, ok = f.(*Group); !ok {
		return out, scope.Error(form.Pos(), "Invalid field list: %v", f)
	}

	for i := 0; i < len(fields.body); i += 2 {
		f = fields.body[i]
		var q *Quote
				
		if q, ok = f.(*Quote); !ok {
			return out, scope.Error(form.Pos(), "Invalid field id: %v", f)
		}

		if _, ok = q.form.(*Id); !ok {
			return out, scope.Error(form.Pos(), "Invalid field id: %v", f)
		}
	}

	var fieldOps []Op
	var err error
		
	if fieldOps, err = fields.Compile(nil, nil, scope); err != nil {
		return out, err
	}
	
	return append(out, NewRecordOp(form, fieldOps)), nil
}

func recordLengthImp(stack *Slice, scope *Scope, pos Pos) (error) {
	stack.Push(NewVal(&TInt, NewInt(int64(stack.Pop().data.(Record).Len()))))
	return nil
}

func recordSetImp(stack *Slice, scope *Scope, pos Pos) (error) {
	v, k, r := stack.Pop(), stack.Pop(), stack.Pop()
	stack.Push(NewVal(&TRecord, r.data.(Record).Set(k.data.(string), *v)))
	return nil
}

func (self *Scope) InitData() *Scope {
	self.AddType(&TRecord)

	self.AddMacro("record:", 1, recordImp)

	self.AddMethod("length", []Argument{AType("val", &TRecord)}, []Result{RType(&TInt)}, recordLengthImp)

	self.AddMethod("set",
		[]Argument{AType("record", &TRecord), AType("key", &TId), AType("val", &TAny)},
		[]Result{RType(&TRecord)},
		recordSetImp)

	return self
}
