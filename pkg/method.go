package gfoo

import (
	"strings"
)

type MethodImp = func(stack *Slice, scope *Scope, pos Pos) (error)

type Method struct {
	function *Function
	arguments []Argument
	results []Result
	imp MethodImp
}

func (self *Method) Init(function *Function, imp MethodImp) *Method{
	self.function = function
	self.imp = imp
	return self
}

func (self *Method) Name() string {
	var name strings.Builder
	name.WriteString(self.function.name)
	name.WriteRune('<')

	for i, a := range self.arguments {
		if i > 0 {
			name.WriteRune(' ')
		}

		a.Dump(&name)
	}

	if self.arguments != nil && self.results != nil {
		name.WriteString("; ")
	}

	for i, r := range self.results {
		r.Dump(&name)
		
		if i > 0 {
			name.WriteRune(' ')
		}
	}

	name.WriteRune('>')
	return name.String()
}
