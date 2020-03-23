package gfoo

type TargetType interface {
	ValType
	Call(target Val, thread *Thread, stack *Slice, pos Pos) error
}
