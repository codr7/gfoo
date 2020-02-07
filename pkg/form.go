package gfoo

type Form interface {
	Quote(*GFoo) *Value
}
