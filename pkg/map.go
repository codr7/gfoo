package gfoo

type Map struct {
	OpBase
	body []Op
}

func NewMap(form Form, body []Op) *Map {
	op := new(Map)
	op.OpBase.Init(form)
	op.body = body
	return op
}

func (self *Map) Eval(thread *Thread, registers, stack *Slice) error {
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing value")
	}

	in, err := v.Iter(self.form.Pos())

	if err != nil {
		return err
	}

	var buffer Slice
	
	stack.Push(NewVal(&TIter, Iter(func (thread *Thread, pos Pos) (Val, error) {
		for {			
			if v := buffer.PopFront(); v != nil {
				return *v, nil
			}
			
			v, err := in(thread, pos)
			
			if err != nil {
				return Nil, err
			}
			
			if v == Nil {
				if buffer.Len() > 0 {
					continue
				}

				break
			}

			buffer.Push(v)

			if err = EvalOps(self.body, thread, registers, &buffer); err != nil {
				return Nil, err
			}
		}

		return Nil, nil
	})))

	return nil
}

