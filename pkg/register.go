package gfoo

const REGISTER_COUNT = 64

type Registers = [REGISTER_COUNT]Val

func NewRegisters() []Val {
	return make([]Val, REGISTER_COUNT)
}
