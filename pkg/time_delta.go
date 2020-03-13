package gfoo

import (
	"fmt"
	"io"
)

type TimeDelta struct {
     years, months, days int
}

func Days(n int) TimeDelta {
     return TimeDelta{days: n}
}

func (self TimeDelta) Compare(other TimeDelta) Order {
	if out := CompareInt(self.years, other.years); out != Eq {
		return out
	}

	if out := CompareInt(self.months, other.months); out != Eq {
		return out
	}

	return CompareInt(self.days, other.days)
}

func (self TimeDelta) Dump(out io.Writer) error {
	_, err := fmt.Fprintf(out, "TimeDelta(%v %v %v)", self.years, self.months, self.days) 
	return err
}
