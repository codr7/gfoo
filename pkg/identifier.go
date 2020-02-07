package gfoo

type Identifier struct {
	name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{name: name}
}

func (self *Identifier) Quote() Value {
	return NewValue(&Symbol, self.name)
}
