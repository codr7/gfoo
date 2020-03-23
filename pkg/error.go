package gfoo

import (
	"errors"
	"fmt"
)

var Debug = false

func NewError(pos Pos, spec string, args...interface{}) error {
	for i, a := range args {
		if d, ok := a.(Dumper); ok {
			args[i] = DumpString(d)
		}
	}
	
	msg := fmt.Sprintf("Error in '%v', line %v, column %v: %v ", 
		pos.source, pos.line, pos.column, fmt.Sprintf(spec, args...))

	return errors.New(msg)
}

func Error(pos Pos, spec string, args...interface{}) error {
	err := NewError(pos, spec, args...)
	
	if Debug {
		panic(err.Error())
	}

	return err
}

