package gfoo

import (
	"fmt"
	"io"
	"strings"
)

type Stack struct {
	items []Val
}

func NewStack(items []Val) *Stack {
	return new(Stack).Init(items)
}

func (self *Stack) Init(items []Val) *Stack {
	self.items = items
	return self
}

func (self *Stack) Clear() {
	self.items = nil
}

func (self *Stack) Clone() *Stack {
	out := make([]Val, len(self.items))

	for i, v := range self.items {
		out[i] = v.Clone()
	}

	return NewStack(out)
}

func (self Stack) Compare(other Stack) Order {
	return CompareVals(self.items, other.items)
}

func (self Stack) Dump(out io.Writer) error {
	if _, err := fmt.Fprint(out, "["); err != nil {
		return err
	}

	if err := DumpVals(self.items, out); err != nil {
		return err
	}
	
	if _, err := fmt.Fprint(out, "]"); err != nil {
		return err
	}
	
	return nil
}

func (self Stack) Len() int {
	return len(self.items)
}

func (self Stack) Peek() *Val {
	i := len(self.items)
	
	if i == 0 {
		return nil
	}

	return &self.items[i-1]
}

func (self *Stack) Pop() *Val {
	i := len(self.items)
	
	if i == 0 {
		return nil
	}

	v := self.items[i-1]
	self.items = self.items[:i-1]
	return &v
}

func (self *Stack) PopFront() *Val {
	if len(self.items) == 0 {
		return nil
	}
	
	v := self.items[0]
	self.items = self.items[1:]
	return &v
}

func (self Stack) Print(out io.Writer) error {
	for _, v := range self.items {
		if err := v.Print(out); err != nil {
			return err
		}
	}
	
	return nil
}

func (self *Stack) Push(vals...Val) {
	self.items = append(self.items, vals...)
}

func (self *Stack) Reset() {
	self.items = nil
}

func (self *Stack) String() string {
	var out strings.Builder
	self.Dump(&out)
	return out.String()
}

func (self *Stack) Unquote(scope *Scope, pos Pos) Form {
	out := make([]Form, len(self.items))

	for i, v := range self.items {
		out[i] = v.Unquote(scope, pos)
	}

	return NewStackForm(out, pos)
}
