package gfoo

type Op interface {
	Eval(*GFoo) error
}
