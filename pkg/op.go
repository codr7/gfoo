package gfoo

type Op interface {
	Evaluate(gfoo *GFoo) error
}
