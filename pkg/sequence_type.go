package gfoo

type SequenceType interface {
	ValType
	Iter(val Val, scope *Scope, pos Pos) (Iter, error)
}
