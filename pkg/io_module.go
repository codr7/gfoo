package gfoo

import (
	"bufio"
	"io"
	"os"
)

type IoModule struct {
	Scope
	ARGS Slice
	OUT *bufio.Writer
}

func bufferLengthImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TInt, Int(stack.Pop().data.(*Buffer).Len())))
	return nil
}

func bufferNewImp(scope *Scope, stack *Slice, pos Pos) error {
	stack.Push(NewVal(&TBuffer, new(Buffer)))
	return nil
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

func writeBufferImp(scope *Scope, stack *Slice, pos Pos) error {
	d := stack.Pop().data.(*Buffer)
	w := stack.Pop().data.(io.Writer)
	_, err := d.WriteTo(w)
	return err
}

func writeStringImp(scope *Scope, stack *Slice, pos Pos) error {
	s := stack.Pop().data.(string)
	w := stack.Pop().data.(io.Writer)
	_, err := io.WriteString(w, s)
	return err
}

func (self *IoModule) Init() *Scope {
	self.Scope.Init()
	self.AddType(&TBuffer)
	self.AddType(&TWriter)

	self.AddVal("ARGS", &TSlice, &self.ARGS)
	self.OUT = bufio.NewWriter(os.Stdout)
	self.AddVal("OUT", &TWriter, self.OUT)

	self.AddMethod("length", []Arg{AType("val", &TBuffer)}, []Ret{RType(&TInt)}, bufferLengthImp)
	self.AddMethod("new-buffer", nil, []Ret{RType(&TBuffer)}, bufferNewImp)
	self.AddMethod("slurp", []Arg{AType("path", &TString)}, []Ret{RType(&TBuffer)}, slurpImp)
	self.AddMethod("write", []Arg{AType("out", &TWriter), AType("data", &TBuffer)}, nil, writeBufferImp)
	self.AddMethod("write", []Arg{AType("out", &TWriter), AType("data", &TString)}, nil, writeStringImp)
	return &self.Scope
}
