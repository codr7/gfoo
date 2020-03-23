package gfoo

type SequenceType interface {
	ValType
	Iter(val Val, pos Pos) (Iter, error)
}
