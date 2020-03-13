package gfoo

import (
	"time"
)

var MinTime, MaxTime time.Time

func init() {
	MinTime = time.Date(0, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	MaxTime = time.Date(9999, time.Month(12), 31, 23, 59, 59, 999999999, time.UTC)
}
