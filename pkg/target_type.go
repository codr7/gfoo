package gfoo

type TargetType interface {
	ValType
	Call(target Val, scope *Scope, stack *Slice, pos Pos) error
}
