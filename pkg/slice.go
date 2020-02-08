package gfoo

import (
	"io"
	"fmt"
)

type Slice struct {
	items []Val
}

func NewSlice(items []Val) *Slice {
	s := new(Slice)
	s.items = items
	return s
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

func (self *Slice) Cut(i int) []Val {
	out := make([]Val, len(self.items)-i)
	copy(out, self.items[i:])
	self.items = self.items[:i]
	return out
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

func (self *Slice) Push(dataType Type, data interface{}) {
	self.items = append(self.items, NewVal(dataType, data))
}

func (self *Slice) Unquote(pos Pos) Form {
	out := make([]Form, len(self.items))

	for i, v := range self.items {
		out[i] = v.Unquote(pos)
	}

	return NewGroup(pos, out)
}
