package gfoo

import (
	"math/big"
)

type Int = big.Int

func NewInt(val int64) *Int {
	return big.NewInt(val)
}
