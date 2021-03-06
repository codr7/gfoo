package gfoo

import (
	"bufio"
	"io"
	"os"
)

type IoModule struct {
	Module
	ARGS Stack
	IN *bufio.Reader
	OUT *bufio.Writer
}

func bufferBytesImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	in, err := stack.Pop().Iter(pos)

	if err != nil {
		return err
	}

	stack.Push(NewVal(&TIter, in))
	return nil
}

func bufferLengthImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	stack.Push(NewVal(&TInt, Int(stack.Pop().data.(*Buffer).Len())))
	return nil
}

func bufferNewImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	stack.Push(NewVal(&TBuffer, new(Buffer)))
	return nil
}

func byteToIntImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	v := stack.Pop().data.(Byte)
	stack.Push(NewVal(&TInt, Int(v)))
	return nil
}

func intToByteImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	v := stack.Pop().data.(Int)

	if v < 0 || v > 255 {
		return Error(pos, "Invalid byte value: %v", v)
	}
	
	stack.Push(NewVal(&TByte, Byte(v)))
	return nil
}

func slurpImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
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

func slurpReaderImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	var b Buffer
	b.ReadFrom(stack.Pop().data.(io.Reader))
	stack.Push(NewVal(&TBuffer, &b))
	return nil
}

func writeBufferImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	d := stack.Pop().data.(*Buffer)
	w := stack.Pop().data.(io.Writer)
	_, err := d.WriteTo(w)
	return err
}

func writeStringImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	s := stack.Pop().data.(string)
	w := stack.Pop().data.(io.Writer)
	_, err := io.WriteString(w, s)
	return err
}

func writeByteImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	b := stack.Pop().data.(byte)
	w := stack.Pop().data.(io.Writer)
	_, err := w.Write([]byte{b})
	return err
}

func (self *IoModule) Init() *Module {
	self.Module.Init()
	
	self.AddType(&TByte)
	self.AddType(&TBuffer)
	self.AddType(&TReader)
	self.AddType(&TWriter)

	self.AddVal("ARGS", &TStack, &self.ARGS)

	self.IN = bufio.NewReader(os.Stdin)
	self.AddVal("IN", &TReader, self.IN)

	self.OUT = bufio.NewWriter(os.Stdout)
	self.AddVal("OUT", &TWriter, self.OUT)

	self.AddMethod("bytes", []Arg{AType("val", &TBuffer)}, []Ret{RType(&TIter)}, bufferBytesImp)
	self.AddMethod("length", []Arg{AType("val", &TBuffer)}, []Ret{RType(&TInt)}, bufferLengthImp)
	self.AddMethod("new-buffer", nil, []Ret{RType(&TBuffer)}, bufferNewImp)

	self.AddMethod("to-int", []Arg{AType("val", &TByte)}, []Ret{RType(&TInt)}, byteToIntImp)
	self.AddMethod("to-byte", []Arg{AType("val", &TInt)}, []Ret{RType(&TByte)}, intToByteImp)
	
	self.AddMethod("slurp", []Arg{AType("path", &TString)}, []Ret{RType(&TBuffer)}, slurpImp)
	self.AddMethod("slurp", []Arg{AType("source", &TReader)}, []Ret{RType(&TBuffer)}, slurpReaderImp)
	self.AddMethod("write", []Arg{AType("out", &TWriter), AType("data", &TBuffer)}, nil, writeBufferImp)
	self.AddMethod("write", []Arg{AType("out", &TWriter), AType("data", &TByte)}, nil, writeByteImp)
	self.AddMethod("write", []Arg{AType("out", &TWriter), AType("data", &TString)}, nil, writeStringImp)

	return &self.Module
}
