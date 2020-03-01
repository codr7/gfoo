package gfoo

import (
	"math/big"
)

func NewInt(val int64) *big.Int {
	return big.NewInt(val)
}
