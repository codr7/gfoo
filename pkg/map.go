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

func (self *Map) Eval(scope *Scope, stack *Slice) error {
	v := stack.Pop()

	if v == nil {
		return scope.Error(self.form.Pos(), "Missing value")
	}

	in, err := v.Iter(scope, self.form.Pos())

	if err != nil {
		return err
	}

	var buffer Slice
	
	stack.Push(NewVal(&TIter, Iter(func (scope *Scope, pos Pos) (Val, error) {
		for {			
			if v := buffer.PopFront(); v != nil {
				return *v, nil
			}
			
			v, err := in(scope, pos)
			
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

			if err = scope.EvalOps(self.body, &buffer); err != nil {
				return Nil, err
			}
		}

		return Nil, nil
	})))

	return nil
}

