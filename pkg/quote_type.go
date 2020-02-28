package gfoo

import (
	"io"
)

var TQuote QuoteType

type QuoteType struct {
	ValTypeBase
}

func (_ *QuoteType) Bool(val Val) bool {
	return val.data.(Val).Bool()
}

func (_ *QuoteType) Compare(x, y Val) Order {
	return x.data.(Val).Compare(y.data.(Val))
}

func (_ *QuoteType) Dump(val Val, out io.Writer) error {
	if _, err := io.WriteString(out, "'"); err != nil {
		return err
	}
	
	return val.data.(Val).Dump(out)
}

func (self *QuoteType) Print(val Val, out io.Writer) error {
	return val.data.(Val).Print(out)
}

func (self *QuoteType) Unquote(val Val, scope *Scope, pos Pos) Form {
	return NewLiteral(val.data.(Val), pos)
}