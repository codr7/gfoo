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

	in, err := v.Iterator(scope, self.form.Pos())

	if err != nil {
		return err
	}

	var buffer Slice
	
	stack.Push(NewVal(&TIterator, Iterator(func (scope *Scope, pos Pos) (*Val, error) {
		for {			
			v := buffer.PopFront()

			if v != nil {
				if *v == Nil {
					continue
				}
				
				return v, nil
			}
			
			v, err = in(scope, pos)
			
			if err != nil {
				return nil, err
			}
			
			if v == nil {
				break
			}
			
			if *v == Nil {
				continue
			}

			scope.val.Push(*v)
			defer scope.val.Pop()

			if err = scope.EvalOps(self.body, &buffer); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})))

	return nil
}

