package gfoo

type Function struct {
	name string
	methods []*Method
}

func NewFunction(name string) *Function {
	return new(Function).Init(name)
}

func (self *Function) Init(name string) *Function {
	self.name = name
	return self
}

func (self *Function) NewMethod(args []Arg, rets []Ret, imp MethodImp, scope *Scope) *Method {
	m := new(Method).Init(self, args, rets, imp, scope)
	self.AddMethod(m)
	return m
}

func (self *Function) AddMethod(method *Method) {
	method.index = len(self.methods)
	self.methods = append(self.methods, method)
}

func (self *Function) RemoveMethod(method *Method) {
	if method.index == -1 {
		panic("Method not added")
	}
	
	if len(self.methods) > method.index {
		self.methods = self.methods[:method.index]
	}
	
	method.index = -1
}

func (self *Function) Call(scope *Scope, stack *Slice, pos Pos) error {
	for i := len(self.methods)-1; i >= 0; i-- {
		if m := self.methods[i]; m.Applicable(stack) {
			return m.Call(stack, pos)
		}
	}
	
	return scope.Error(pos, "Function not applicable: %v %v", self.name, stack)
}


