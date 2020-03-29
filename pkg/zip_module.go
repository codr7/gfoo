package gfoo

import (
	"archive/zip"
)

type ZipModule struct {
	Module
}

func zipAddImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	p := stack.Pop().data.(string)
	w, err := stack.Pop().data.(*zip.Writer).Create(p)

	if err != nil {
		return err
	}

	stack.Push(NewVal(&TWriter, w))
	return nil
}

func zipCloseImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	return stack.Pop().data.(*zip.Writer).Close()
}

func zipWriterNewImp(thread *Thread, registers []Val, stack *Stack, pos Pos) error {
	out := stack.Pop().data.(*Buffer)
	stack.Push(NewVal(&TZipWriter, zip.NewWriter(out)))
	return nil
}

func (self *ZipModule) Init() *Module {
	self.Module.Init()
	
	self.AddType(&TZipWriter)

	self.AddMethod("add",
		[]Arg{AType("zip", &TZipWriter), AType("path", &TString)},
		[]Ret{RType(&TWriter)},
		zipAddImp)

	self.AddMethod("close", []Arg{AType("zip", &TZipWriter)}, nil, zipCloseImp)
	self.AddMethod("new-writer", []Arg{AType("out", &TBuffer)}, []Ret{RType(&TZipWriter)}, zipWriterNewImp)

	return &self.Module
}
