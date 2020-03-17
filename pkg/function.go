package gfoo

import (
	"strings"
)

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

func (self *Function) AddMethod(methods...*Method) {
	for _, m := range methods {
		m.indexes[self] = len(self.methods)
		self.methods = append(self.methods, m)
	}
}

func (self *Function) Call(scope *Scope, stack *Slice, pos Pos) error {
	for i := len(self.methods)-1; i >= 0; i-- {
		if m := self.methods[i]; m.Applicable(stack) {
			return m.Call(scope, stack, pos)
		}
	}
	
	return scope.Error(pos, "Function not applicable: %v %v", self.name, stack)
}

func (self *Function) MethodName(args []Arg, rets []Ret) string {
	var out strings.Builder
	out.WriteString(self.name)
	out.WriteRune('<')

	for i, a := range args {
		if i > 0 {
			out.WriteRune(' ')
		}

		a.Dump(&out)
	}

	out.WriteRune(';')

	for _, r := range rets {
		out.WriteRune(' ')
		r.Dump(&out)
	}

	out.WriteRune('>')
	return out.String()
}

func (self *Function) NewMethod(args []Arg, rets []Ret, imp MethodImp) *Method {
	m := new(Method).Init(self.MethodName(args, rets), args, rets, imp)
	self.AddMethod(m)
	return m
}

func (self *Function) RemoveMethod(method *Method) {
	index, ok := method.indexes[self]
	
	if !ok {
		panic("Method not added")
	}

	if len(self.methods) > index {
		for i := index; i < len(self.methods); i++ {
			delete(self.methods[i].indexes, self)
		}
		
		self.methods = self.methods[:index]
	}
}

