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

func (self *Function) AddMethod(imp MethodImp) *Method {
	m := new(Method).Init(self, imp)
	self.methods = append(self.methods, m)
	return m
}

