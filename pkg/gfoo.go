package gfoo

import (
	"sync"
)

type GFoo struct {
	symbols sync.Map
}

func New() *GFoo {
	gfoo := new(GFoo)
	return gfoo
}

func (gfoo *GFoo) Symbol(name string) *Value {
	found, _ := gfoo.symbols.Load(name)

	if found != nil {
		return found.(*Value)
	}
	
	s := NewValue(&Symbol, name)
	gfoo.symbols.Store(name, s)
	return s
}
