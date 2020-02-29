package gfoo

import (
	"fmt"
	"io"
)

type Record struct {
	fields Tree
}

func NewRecord() Record {
	var r Record

	r.fields.Init(func (x, y interface{}) Order {
		return CompareString(x.(string), y.(string))
	})

	return r
}

func (self Record) Compare(other Record) Order {
	return self.fields.Compare(&other.fields, func (x, y *TreeNode) Order {
		if out := CompareString(x.key.(string), y.key.(string)); out != Eq {
			return out
		}
		
		return x.values[0].(Val).Compare(y.values[0].(Val))
	})
}

func (self Record) Dump(out io.Writer) error {
	if _, err := io.WriteString(out, "("); err != nil {
		return err
	}

	first := true
	
	if err := self.fields.ForEach(func (k, v interface{}) error {
		if first {
			first = false
		} else {
			if _, err := io.WriteString(out, " "); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintf(out, "%v ", k.(string)); err != nil {
			return err
		}
		
		return v.(Val).Dump(out)
	}); err != nil {
		return err
	}

	if _, err := io.WriteString(out, ")"); err != nil {
		return err
	}

	return nil
}

func (self Record) Get(key string, val Val) Val {
	if found := self.fields.Find(key); found != nil {
		return found[0].(Val)
	}

	return val
}

func (self Record) Len() uint64 {
	return self.fields.Len()
}

func (self Record) Set(key string, val Val) Record {
	self.fields = self.fields.Update(key, val)
	return self
}
