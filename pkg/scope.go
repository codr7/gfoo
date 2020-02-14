package gfoo

import (
	"bufio"
	"io"
)

type Bindings = map[string]Binding

type Scope struct {
	thread *Thread
	bindings Bindings
	Debug bool
}

func (self *Scope) Init(thread *Thread) *Scope {
	self.thread = thread
	self.bindings = make(Bindings)
	return self
}

func (self *Scope) Compile(in []Form, out []Op) ([]Op, error) {
	var err error
	var inForms Forms
	inForms.Init(in)
	
	for f := inForms.Pop(); f != nil; f = inForms.Pop() {
		if out, err = f.Compile(&inForms, out, self); err != nil {
			return out, err
		}
	}
	
	return out, nil
}

func (self *Scope) Copy(out *Scope, child bool) {
	if child {
		out.Debug = self.Debug
	}
	
	for k, b := range self.bindings {
		if !child && b.scope == self {
			b.scope = out
		}
		
		out.bindings[k] = b
	}
}

func (self *Scope) Clone(child bool) *Scope {
	out := new(Scope).Init(self.thread)
	self.Copy(out, child)
	return out
}

func (self *Scope) Evaluate(ops []Op, stack *Slice) error {
	for _, o := range ops {
		if err := o.Evaluate(stack, self); err != nil {
			return err
		}
	}
	
	return nil
}

func (self *Scope) Get(key string) *Binding {
	if found, ok := self.bindings[key]; ok {
		return &found
	}

	return nil
}

func (self *Scope) Parse(in *bufio.Reader, pos *Pos, out []Form) ([]Form, error) {
	var f Form
	var err error
	
	for {
		if err = skipSpace(in, pos); err == nil {
			f, err = self.parseForm(in, pos)
		}

		if err == io.EOF {
			break
		}

		if err != nil {			
			return out, err
		}

		out = append(out, f)
	}
	
	return out, nil
}

func (self *Scope) Set(key string, val Val) {
	self.bindings[key] = NewBinding(self, val)
}
