package gfoo

import (
	"bufio"
	"io"
	"os"
)

type IoModule struct {
	Scope
	OUT *bufio.Writer
}

func bufferNewImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TBuffer, new(Buffer)))
	return nil
}

func bufferWriteImp(scope *Scope, stack *Slice, pos Pos) error {
	out := stack.Pop().data.(io.Writer)
	_, err := stack.Pop().data.(*Buffer).WriteTo(out)
	return err
}

func slurpImp(scope *Scope, stack *Slice, pos Pos) error {
	p := stack.Pop().data.(string)
	f, err := os.Open(p)
	
	if err != nil {
		return err
	}

	var b Buffer
	b.ReadFrom(bufio.NewReader(f))
	stack.Push(NewVal(&TBuffer, &b))
	return nil
}

func (self *IoModule) Init() *Scope {
	self.Scope.Init()
	self.AddType(&TBuffer)
	self.AddType(&TWriter)

	self.OUT = bufio.NewWriter(os.Stdout)
	self.AddVal("OUT", &TWriter, self.OUT)

	self.AddMethod("new-buffer", nil, []Ret{RType(&TBuffer)}, bufferNewImp)
	self.AddMethod("write", []Arg{AType("data", &TBuffer), AType("out", &TWriter)}, nil, bufferWriteImp)
	self.AddMethod("slurp", []Arg{AType("path", &TString)}, []Ret{RType(&TBuffer)}, slurpImp)
	return &self.Scope
}
