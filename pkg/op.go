package gfoo

type Op interface {
	Evaluate(*GFoo) error
}
