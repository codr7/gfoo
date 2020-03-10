package gfoo

import (
	"io"
)

type Pair struct {
	left, right Val
}

func NewPair(left, right Val) Pair {
	return Pair{left, right}
}

func (self Pair) Compare(other Pair) Order {
	if out := self.left.Compare(other.left); out != Eq {
		return out
	}

	return self.right.Compare(other.right)
}

func (self Pair) Dump(out io.Writer) error {
	if err := self.left.Dump(out); err != nil {
		return err
	}

	if _, err := io.WriteString(out, " "); err != nil {
		return err
	}
	
	if err := self.right.Dump(out); err != nil {
		return err
	}

	if _, err := io.WriteString(out, ","); err != nil {
		return err
	}

	return nil
}

func (self Pair) Print(out io.Writer) error {
	if err := self.left.Print(out); err != nil {
		return err
	}

	if err := self.right.Print(out); err != nil {
		return err
	}

	return nil
}
