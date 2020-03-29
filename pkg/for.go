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

func (self *For) Eval(thread *Thread, registers []Val, stack *Stack) error {
	v := stack.Pop()

	if v == nil {
		return Error(self.form.Pos(), "Missing source")
	}

	in, err := v.Iter(self.form.Pos())

	if err != nil {
		return err
	}

	if self.id != -1 {
		registers[self.id] = Nil
	}
	
	return in.For(func(v Val, thread *Thread, pos Pos) error {
		if self.id == -1 {
			stack.Push(v)
		} else {
			registers[self.id] = v
		}
		
		return EvalOps(self.body, thread, registers, stack)
	}, thread, self.form.Pos())
}
