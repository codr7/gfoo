package gfoo

import (
	"strings"
	"unsafe"
)

type Order = int

const (
	Lt = Order(-1)
	Eq = Order(0)
	Gt = Order(1)
)

func CompareByte(x, y uint8) Order {
	if x < y {
		return Lt
	}

	if x > y {
		return Gt
	}

	return Eq
}

func CompareInt(x, y int) Order {
	if x < y {
		return Lt
	}

	if x > y {
		return Gt
	}

	return Eq
}

func CompareInt64(x, y Int) Order {
	if x < y {
		return Lt
	}

	if x > y {
		return Gt
	}

	return Eq
}

func ComparePointer(x, y unsafe.Pointer) Order {
	xp, yp := uintptr(x), uintptr(y)
	
	if xp < yp {
		return Lt
	}
	
	if xp > yp {
		return Gt
	}
	
	return Eq
}

func CompareRune(x, y rune) Order {
	if x < y {
		return Lt
	}

	if x > y {
		return Gt
	}

	return Eq
}

func CompareString(x, y string) Order {
	return Order(strings.Compare(x, y))
}

func MinInt(x, y int) int {
	if y < x {
		return y
	}

	return x
}
