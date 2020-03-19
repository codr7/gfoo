package gfoo

import (
	"bufio"
	"io"
	"path"
	"os"
	"strings"
)

type Bindings = map[string]Binding

type Scope struct {
	Debug bool
	thread *Thread
	loadPath string
	bindings Bindings
	methods []*Method
	val Slice
}

func NewScope() *Scope {
	return new(Scope).Init()
}

func (self *Scope) Init() *Scope {
	self.bindings = make(Bindings)
	return self
}

func (self *Scope) AddFunction(name string) *Function {
	f := NewFunction(name)
	self.AddVal(name, &TFunction, f)
	return f
}

func (self *Scope) AddMacro(name string, argCount int, imp MacroImp) {
	self.AddVal(name, &TMacro, NewMacro(name, argCount, imp))
}

func (self *Scope) AddMethod(name string, args []Arg, rets []Ret, imp MethodImp) {
	var f *Function
	b := self.Get(name)
	
	if b == nil || b.val == Nil {
		f = self.AddFunction(name)
	} else {
		f = b.val.data.(*Function)
	}

	m := f.NewMethod(args, rets, imp)
	self.AddVal(m.name, &TMethod, m)
	self.methods = append(self.methods, m)
}

func (self *Scope) AddModule(name string, module *Module) {
	self.AddVal(name, &TModule, module)
}

func (self *Scope) AddType(val Type) {
	self.AddVal(val.Name(), &TMeta, val)
}

func (self *Scope) AddVal(name string, dataType ValType, data interface{}) {
	self.Set(name, NewVal(dataType, data))
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
	out.loadPath = self.loadPath
	out.val = self.val
	
	for k, b := range self.bindings {
		out.bindings[k] = b
	}
}

func (self *Scope) Extend(source *Scope) *Scope {
	for k, b := range source.bindings {
		self.bindings[k] = b
	}

	self.methods = append(self.methods, source.methods...)
	self.val.Push(source.val.items...)
	return self
}

func (self *Scope) Clone() *Scope {
	out := new(Scope).Init()
	self.Copy(out)
	return out
}

func (self *Scope) Eval(source string, stack *Slice) error {
	in := bufio.NewReader(strings.NewReader(source))
	pos := NewPos("n/a")
	var forms []Form
	var err error
	
	if forms, err = self.Parse(in, nil, &pos); err != nil {
		return err
	}
	
	var ops []Op
	
	if ops, err = self.Compile(forms, nil); err != nil {
		return err
	}

	if err = self.EvalOps(ops, stack); err != nil {
		return err
	}

	return nil
}

func (self *Scope) EvalOps(ops []Op, stack *Slice) error {
	for _, o := range ops {
		if err := o.Eval(self, stack); err != nil {
			return err
		}
	}
	
	return nil
}

func (self *Scope) EvalForm(in *Forms, stack *Slice) error {
	f := in.Pop()

	if f == nil {
		return nil
	}
	
	ops, err := f.Compile(in, nil, self)
	
	if err != nil {
		return err
	}
	
	if err = self.EvalOps(ops, stack); err != nil {
		return err
	}

	return nil
}

func (self *Scope) Get(key string) *Binding {
	if found, ok := self.bindings[key]; ok {
		return &found
	}

	return nil
}

func (self *Scope) Include(filePath string, action func([]Form) error) (error) {
	var file *os.File
	var err error

	prevLoadPath := self.loadPath
	filePath = path.Join(self.loadPath, filePath)
	self.loadPath = path.Dir(filePath)
		
	defer func() {
		self.loadPath = prevLoadPath
	}()
	
	if file, err = os.Open(filePath); err != nil {
		return err
	}

	in := bufio.NewReader(file)
	pos := NewPos(filePath)
	var forms []Form
	
	if forms, err = self.Parse(in, nil, &pos); err != nil {
		return err
	}

	return action(forms)
}

func (self *Scope) Load(filePath string, stack *Slice) error {
	return self.Include(filePath, func(forms []Form) error {
		var ops []Op
		var err error
		
		if ops, err = self.Clone().Compile(forms, nil); err != nil {
			return err
		}
		
		if err = self.EvalOps(ops, stack); err != nil {
			return err
		}
		
		return nil
	})
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

func (self *Scope) Use(source Val, names []string, pos Pos) error {
	useAll := false
	
	if len(names) == 0 {
		names = source.Keys()
		useAll = true
	}

	for _, n := range names {
		v, err := source.Get(n, self, pos)

		if err != nil {
			return err
		}

		if found := self.Get(n); found != nil {
			if v.dataType == &TFunction && found.val.dataType == &TFunction {
				v.data.(*Function).AddMethod(found.val.data.(*Function).methods...)
			} else if !useAll {
				return self.Error(pos, "Duplicate identifier: %v", n)
			}
		}

		self.Set(n, v)
	}

	return nil
}
