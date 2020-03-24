package gfoo

type For struct {
	OpBase
	id int
	body []Op
}

func NewFor(form Form, id int, body []Op) *For {
	op := new(For)
	op.OpBase.Init(form)
	op.id = id
	op.body = body
	return op
}

func (self *For) Eval(thread *Thread, registers, stack *Slice) error {
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing source")
	}

	in, err := v.Iter(self.form.Pos())

	if err != nil {
		return err
	}

	if self.id != -1 {
		if registers.Len() <= self.id {
			registers.Push(Nil)
		} else {
			registers.items[self.id] = Nil
		}
	}
	
	return in.For(func(v Val, thread *Thread, pos Pos) error {
		if self.id == -1 {
			stack.Push(v)
		} else {
			registers.items[self.id] = v
		}
		
		if err = EvalOps(self.body, thread, registers, stack); err != nil && err != &Break {
			return err
		}
			
		return nil
	}, thread, self.form.Pos())
}
