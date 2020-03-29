package gfoo

type IterModule struct {
	Module
}

func iterChainImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	y := stack.Pop().data.(Iter)
	x := stack.Pop().data.(Iter)

	stack.Push(NewVal(&TIter, Iter(func(thread *Thread, pos Pos) (Val, error) {
		if x != nil {
			v, err := x(thread, pos)

			if err != nil {
				return Nil, err
			}

			if v != Nil {
				return v, nil
			}

			x = nil
		}

		if y != nil {
			v, err := y(thread, pos)

			if err != nil {
				return Nil, err
			}

			if v != Nil {
				return v, nil
			}

			y = nil
		}
		
		return Nil, nil
	})))
	
	return nil
}

func iterNextImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	in := stack.Pop().data.(Iter)
	
	if v, err := in(thread, pos); err != nil {
		return err
	} else {
		stack.Push(v)
	}
			
	return nil
}

func iterValueImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	v := *stack.Pop()

	stack.Push(NewVal(&TIter, Iter(func(thread *Thread, pos Pos) (Val, error) {
		return v, nil
	})))
	
	return nil
}

func (self *IterModule) Init() *Module {
	self.Module.Init()
	
	self.AddMethod("~", []Arg{AType("x", &TIter), AType("y", &TIter)}, []Ret{RType(&TIter)}, iterChainImp)
	self.AddMethod("next", []Arg{AType("in", &TIter)}, []Ret{RType(&TOption)}, iterNextImp)
	self.AddMethod("value", []Arg{AType("val", &TAny)}, []Ret{RType(&TIter)}, iterValueImp)
	
	return &self.Module
}
