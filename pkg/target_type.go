package gfoo

type TargetType interface {
	ValType
	Call(target Val, thread *Thread, stack *Stack, pos Pos) error
}
