package gfoo

type Identifier struct {
	name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{name: name}
}

func (id *Identifier) Quote(gfoo *GFoo) Value {
	return NewValue(&Symbol, id.name)
}
