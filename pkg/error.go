package gfoo

import (
	"errors"
	"fmt"
)

func NewError(pos Pos, spec string, args...interface{}) error {
	msg := fmt.Sprintf("Error in '%v', line %v, column %v: %v ", 
		pos.source, pos.line, pos.column, fmt.Sprintf(spec, args...))

	return errors.New(msg)
}
