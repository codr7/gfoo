package gfoo

import (
	"strings"
)

type MethodImp = func(stack *Slice, scope *Scope, pos Pos) error

type Method struct {
	function *Function
	arguments []Argument
	results []Result
	imp MethodImp
	scope *Scope
}

func (self *Method) Init(
	function *Function,
	arguments []Argument,
	results []Result,
	imp MethodImp,
	scope *Scope) *Method{
	self.function = function
	self.arguments = arguments
	self.results = results
	self.imp = imp
	self.scope = scope.Clone()
	return self
}

func (self *Method) Applicable(stack *Slice) bool {
	sl, al := stack.Len(), len(self.arguments)
	
	if sl < al {
		return false
	}

	s := stack.items[sl-al:]
	si := 0
	
	for _, a := range self.arguments {
		if !a.Match(s, si) {
			return false
		}

		si++
	}
	
	return true
}

func (self *Method) Call(stack *Slice, pos Pos) error {
	if stack.Len() < len(self.arguments) {
		return self.scope.Error(pos, "Method not applicable: %v %v", self.Name(), stack)
	}
	
	return self.imp(stack, self.scope.Clone(), pos)
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

	if self.results != nil {
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
