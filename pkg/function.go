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

func (self *Function) AddMethod(arguments, results []Argument, imp MethodImp, scope *Scope) *Method {
	m := new(Method).Init(self, arguments, results, imp, scope)
	self.methods = append(self.methods, m)
	return m
}

func (self *Function) Call(scope *Scope, stack *Slice, pos Pos) error {
	return scope.Error(pos, "Function not applicable: %v %v", self.name, stack)
}


