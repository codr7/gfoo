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
	scope *Scope
}

func (self *Method) Init(function *Function, imp MethodImp, scope *Scope) *Method{
	self.function = function
	self.imp = imp
	self.scope = scope.Clone()
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

func (self *Method) Call(stack *Slice, pos Pos) error {
	if sl, ac := stack.Len(), len(self.arguments); sl < ac {
		self.scope.Error(pos, "Not enough arguments: %v (%v)", sl, ac)
	}
	
	return self.imp(stack, self.scope.Clone(), pos)
}
