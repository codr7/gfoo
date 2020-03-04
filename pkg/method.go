package gfoo

import (
	"strings"
)

type MethodImp = func(stack *Slice, scope *Scope, pos Pos) error

type Method struct {
	function *Function
	index int
	args []Arg
	rets []Ret
	imp MethodImp
	scope *Scope
}

func (self *Method) Init(
	function *Function,
	args []Arg,
	rets []Ret,
	imp MethodImp,
	scope *Scope) *Method{
	self.function = function
	self.index = -1
	self.args = args
	self.rets = rets
	self.imp = imp
	self.scope = scope
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

func (self *Method) Call(stack *Slice, pos Pos) error {	
	var in Slice
	argCount := len(self.args)

	if argCount > 0 {
		in.items = make([]Val, argCount)
		copy(in.items, stack.items[stack.Len()-argCount:])
	}
	
	if err := self.imp(stack, self.scope.Clone(), pos); err != nil {
		return err
	}

	retCount := len(self.rets)
	var out Slice
	out.items = stack.items[stack.Len()-retCount:]
		
	for i := 0; i < retCount; i++ {
		if !self.rets[i].Match(in.items, out.items, i) {
			return self.scope.Error(pos, "Invalid method result: %v %v", self.Name(), out)
		}
	}

	return nil
}

func (self *Method) Name() string {
	var name strings.Builder
	name.WriteString(self.function.name)
	name.WriteRune('<')

	for i, a := range self.args {
		if i > 0 {
			name.WriteRune(' ')
		}

		a.Dump(&name)
	}

	if self.rets != nil {
		name.WriteString("; ")
	}

	for i, r := range self.rets {
		r.Dump(&name)
		
		if i > 0 {
			name.WriteRune(' ')
		}
	}

	name.WriteRune('>')
	return name.String()
}
