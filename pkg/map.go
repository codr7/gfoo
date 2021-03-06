package gfoo

type Map struct {
	OpBase
	body []Op
	id int
}

func NewMap(form Form, id int, body []Op) *Map {
	op := new(Map)
	op.OpBase.Init(form)
	op.body = body
	op.id = id
	return op
}

func (self *Map) Eval(thread *Thread, registers []Val, stack *Stack) error {
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing value")
	}

	in, err := v.Iter(self.form.Pos())

	if err != nil {
		return err
	}
	
	if self.id != -1 {
		registers[self.id] = Nil
	}

	var buffer Stack

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
				break
			}
			
			if self.id == -1 {
				buffer.Push(v)
			} else {
				registers[self.id] = v
			}

			if err = EvalOps(self.body, thread, registers, &buffer); err != nil {
				return Nil, err
			}
		}

		return Nil, nil
	})))

	return nil
}

