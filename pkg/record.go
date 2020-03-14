package gfoo

import (
	"fmt"
	"io"
)

type Record struct {
	fields []RecordField
}

type RecordField struct {
	key string
	val Val
}

func NewRecord() *Record {
	return new(Record)
}

func (self *Record) Clone() *Record {
	out := NewRecord()
	out.fields = make([]RecordField, self.Len())
	copy(out.fields, self.fields)
	return out
}

func (self *Record) Compare(other *Record) Order {
	for i, x := range self.fields {
		y := other.fields[i]
		
		if out := CompareString(x.key, y.key); out != Eq {
			return out
		}
		
		if out := x.val.Compare(y.val); out != Eq {
			return out
		}
	}

	return Eq
}

func (self *Record) Dump(out io.Writer) error {
	if _, err := io.WriteString(out, "("); err != nil {
		return err
	}

	for i, f := range self.fields {
		if i > 0 {
			if _, err := io.WriteString(out, " "); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintf(out, "%v ", f.key); err != nil {
			return err
		}
		
		f.val.Dump(out)
	}

	if _, err := io.WriteString(out, ")"); err != nil {
		return err
	}

	return nil
}

func (self *Record) Find(key string) (int, bool) {
	min, max := 0, self.Len()

	for min < max {
		i := (min+max) / 2
		f := self.fields[i]

		switch CompareString(key, f.key) {
		case Lt:
			max = i
		case Gt:
			min = i+1
		default:
			return i, true
		}
	}

	return min, false
}

func (self *Record) Get(key string, missingVal Val) Val {
	if i, ok := self.Find(key); ok {
		return self.fields[i].val
	}

	return missingVal
}

func (self *Record) Insert(i int, key string, val Val) {
	f := RecordField{key: key, val: val}
	l := self.Len()
	
	if i == l {
		self.fields = append(self.fields, f)
	} else if i == l-1 {
		self.fields = append(self.fields[:i], f, self.fields[i])
	} else {
		prev := self.fields
		self.fields = make([]RecordField, l+1)
		copy(self.fields, prev[:i])
		self.fields[i] = f
		copy(self.fields[i+1:], prev[i:])
	}
}

func (self *Record) Len() int {
	return len(self.fields)
}

func (self *Record) Merge(source *Record) {
	xi, yi := 0, 0
	xl, yl := self.Len(), source.Len()
	var out []RecordField
	
	for xi < xl && yi < yl {
		x, y := self.fields[xi], source.fields[yi]
		
		switch CompareString(x.key, y.key) {
		case Lt:
			out = append(out, x)
			xi++
		case Gt:
			out = append(out, y)
			yi++
		default:
			out = append(out, x)
			xi++
			yi++
		}
	}

	for xi < xl {
		out = append(out, self.fields[xi])
		xi++
	}

	for yi < yl {
		out = append(out, source.fields[yi])
		yi++
	}

	self.fields = out
}

func (self *Record) Set(key string, val Val) {
	if i, ok := self.Find(key); ok {
		self.fields[i].val = val
	} else {
		self.Insert(i, key, val)
	}
}
