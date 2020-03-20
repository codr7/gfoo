package gfoo

type IterModule struct {
	Module
}

func iterChainImp(scope *Scope, stack *Slice, pos Pos) error {
	y := stack.Pop().data.(Iter)
	x := stack.Pop().data.(Iter)

	stack.Push(NewVal(&TIter, Iter(func(scope *Scope, pos Pos) (Val, error) {
		if x != nil {
			v, err := x(scope, pos)

			if err != nil {
				return Nil, err
			}

			if v != Nil {
				return v, nil
			}

			x = nil
		}

		if y != nil {
			v, err := y(scope, pos)

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

func iterNextImp(scope *Scope, stack *Slice, pos Pos) error {
	in := stack.Pop().data.(Iter)
	
	if v, err := in(scope, pos); err != nil {
		return err
	} else {
		stack.Push(v)
	}
			
	return nil
}

func iterValueImp(scope *Scope, stack *Slice, pos Pos) error {
	v := *stack.Pop()
	stack.Push(NewVal(&TIter, Iter(func(scope *Scope, pos Pos) (Val, error) { return v, nil })))
	return nil
}

func (self *IterModule) Init() *Module {
	self.Module.Init()
	
	self.AddMethod("~", []Arg{AType("x", &TIter), AType("y", &TIter)}, []Ret{RType(&TIter)}, iterChainImp)
	self.AddMethod("next", []Arg{AType("in", &TIter)}, []Ret{RType(&TOption)}, iterNextImp)
	self.AddMethod("value", []Arg{AType("val", &TAny)}, []Ret{RType(&TIter)}, iterValueImp)
	
	return &self.Module
}
