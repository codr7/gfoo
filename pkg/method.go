package gfoo

type MethodImp = func(thread *Thread, registers, stack *Slice, pos Pos) error

type Method struct {
	indexes map[*Function]int
	name string
	args []Arg
	rets []Ret
	imp MethodImp
	registers Slice
}

func (self *Method) Init(
	name string,
	args []Arg,
	rets []Ret,
	imp MethodImp) *Method{
	self.indexes = make(map[*Function]int)
	self.name = name
	self.args = args
	self.rets = rets
	self.imp = imp
	return self
}

func (self *Method) Applicable(stack *Slice) bool {
	sl, al := stack.Len(), len(self.args)
	
	if sl < al {
		return false
	}

	s := stack.items[sl-al:]
	si := 0
	
	for _, a := range self.args {
		if !a.Match(s, si) {
			return false
		}

		si++
	}
	
	return true
}

func (self *Method) Call(thread *Thread, stack *Slice, pos Pos) error {	
	var in []Val
	argCount := len(self.args)

	if argCount > 0 {
		in = make([]Val, argCount)
		copy(in, stack.items[stack.Len()-argCount:])
	}

	if err := self.imp(thread, &self.registers, stack, pos); err != nil {
		return err
	}

	retCount := len(self.rets)

	if stack.Len() < retCount {
		return Error(pos, "Missing method result: %v %v", self.name, stack)
	}
	
	offs := stack.Len()-retCount
	
	for i := offs; i < stack.Len(); i++ {
		if !self.rets[i-offs].Match(in, stack.items, i) {
			return Error(pos, "Invalid method result: %v %v", self.name, stack)
		}
	}

	return nil
}
