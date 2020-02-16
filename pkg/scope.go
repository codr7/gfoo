package gfoo

import (
	"bufio"
	"fmt"
	"io"
)

type Bindings = map[string]Binding

type Scope struct {
	Debug bool
	thread *Thread
	bindings Bindings
}

func (self *Scope) Init() *Scope {
	self.bindings = make(Bindings)
	return self
}

func (self *Scope) AddConst(name string, dataType Type, data interface{}) bool {
	if found := self.Get(name); found != nil {
		return false
	}

	self.Set(name, NewVal(dataType, data))
	return true
}

func (self *Scope) AddMacro(name string, argCount int, imp MacroImp) bool {
	return self.AddConst(name, &TMacro, NewMacro(name, argCount, imp))
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

func (self *Scope) Copy(out *Scope) {
	out.Debug = self.Debug
	out.thread = self.thread
	
	for k, b := range self.bindings {
		out.bindings[k] = b
	}
}

func (self *Scope) Clone() *Scope {
	out := new(Scope).Init()
	self.Copy(out)
	return out
}

func (self *Scope) Evaluate(ops []Op, stack *Slice) error {
	for _, o := range ops {
		if err := o.Evaluate(self, stack); err != nil {
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

func (self *Scope) Parse(in *bufio.Reader, out []Form, pos *Pos) ([]Form, error) {
	var f Form
	var err error
	
	for {
		if err = SkipSpace(in, pos); err == nil {
			f, err = self.ParseForm(in, pos)
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

func (self *Scope) Unique(key string) string {
	if b, ok := self.bindings[key]; ok && b.val.data != nil {
		return b.val.data.(string)
	}

	out := fmt.Sprintf("%v%v", key, len(self.bindings))
	self.Set(key, NewVal(&TString, out))
	return out
}
