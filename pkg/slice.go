package gfoo

import (
	"fmt"
	"io"
	"strings"
)

type Slice struct {
	items []Val
}

func NewSlice(items []Val) *Slice {
	return new(Slice).Init(items)
}

func (self *Slice) Init(items []Val) *Slice {
	self.items = items
	return self
}

func (self *Slice) Clear() {
	self.items = nil
}

func (self *Slice) Clone() *Slice {
	out := make([]Val, len(self.items))

	for i, v := range self.items {
		out[i] = v.Clone()
	}

	return NewSlice(out)
}

func (self Slice) Compare(other Slice) Order {
	xl, yl := len(self.items), len(other.items)
	
	for i := 0; i < MinInt(xl, yl); i++ {
		if o := self.items[i].Compare(other.items[i]); o != Eq {
			return o
		}		
	}
	
	return CompareInt(xl, yl)
}

func (self Slice) Dump(out io.Writer) error {
	if _, err := fmt.Fprint(out, "["); err != nil {
		return err
	}

	for i, v := range self.items {
		if i > 0 {
			if _, err := fmt.Fprint(out, " "); err != nil {
				return err
			}
		}
		
		if err := v.Dump(out); err != nil {
			return err
		}
	}
	
	if _, err := fmt.Fprint(out, "]"); err != nil {
		return err
	}
	
	return nil
}

func (self Slice) Len() int {
	return len(self.items)
}

func (self Slice) Peek() *Val {
	i := len(self.items)
	
	if i == 0 {
		return nil
	}

	return &self.items[i-1]
}

func (self *Slice) Pop() *Val {
	i := len(self.items)
	
	if i == 0 {
		return nil
	}

	v := self.items[i-1]
	self.items = self.items[:i-1]
	return &v
}

func (self Slice) Print(out io.Writer) error {
	for i, v := range self.items {
		if i > 0 {
			if _, err := fmt.Fprint(out, " "); err != nil {
				return err
			}
		}
		
		if err := v.Print(out); err != nil {
			return err
		}
	}
	
	return nil
}

func (self *Slice) Push(vals...Val) {
	self.items = append(self.items, vals...)
}

func (self *Slice) Reset() {
	self.items = nil
}

func (self *Slice) String() string {
	var out strings.Builder
	self.Dump(&out)
	return out.String()
}

func (self *Slice) Unquote(scope *Scope, pos Pos) Form {
	out := make([]Form, len(self.items))

	for i, v := range self.items {
		out[i] = v.Unquote(scope, pos)
	}

	return NewGroup(out, pos)
}
