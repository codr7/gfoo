package gfoo

import (
	"bufio"
	"io"
	"path"
	"os"
	"strings"
)

type Scope struct {
	parent *Scope
	loadPath string
	bindings map[string]Binding
	registers map[string]int
	methods []*Method
}

func NewScope(parent *Scope) *Scope {
	return new(Scope).Init(parent)
}

func (self *Scope) Init(parent *Scope) *Scope {
	self.parent = parent
	self.bindings = make(map[string]Binding)
	self.registers = make(map[string]int)

	if parent != nil {
		self.loadPath = parent.loadPath
	
		for k, v := range parent.registers {
			self.registers[k] = v
		}
	}
	
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

func (self *Scope) AddMethod(name string, args []Arg, rets []Ret, imp MethodImp) *Method {
	var f *Function
	b := self.Get(name)
	
	if b == nil || b.val == Undefined {
		f = self.AddFunction(name)
	} else {
		f = b.val.data.(*Function)
	}

	m := f.NewMethod(args, rets, imp)
	self.AddVal(m.name, &TMethod, m)
	self.methods = append(self.methods, m)
	return m
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

func (self *Scope) Extend(source *Scope) *Scope {
	for k, b := range self.bindings {
		if b.val == Undefined {
			self.bindings[k] = source.bindings[k]
		}
	}

	self.methods = append(self.methods, source.methods...)
	return self
}

func (self *Scope) Eval(source string, thread *Thread, registers []Val, stack *Stack) error {
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

	if err = EvalOps(ops, thread, registers, stack); err != nil {
		return err
	}

	return nil
}

func (self *Scope) EvalForm(in *Forms, stack *Stack) error {
	f := in.Pop()

	if f == nil {
		return nil
	}
	
	ops, err := f.Compile(in, nil, NewScope(self))
	
	if err != nil {
		return err
	}

	var registers Registers
	
	if err = EvalOps(ops, nil, registers[:], stack); err != nil {
		return err
	}

	return nil
}

func (self *Scope) Get(key string) *Binding {
	if found, ok := self.bindings[key]; ok {
		return &found
	}

	if self.parent != nil {
		return self.parent.Get(key)
	}
	
	return nil
}

func (self *Scope) Include(filePath string, action func([]Form) error) error {
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

	if c, _, err := in.ReadRune(); err != nil {
		if err != io.EOF {
			return err
		}
	} else if c == '#' {
		if _, err := in.ReadString('\n'); err != nil {
			return err
		}
	} else if err := in.UnreadRune(); err != nil {
		return err
	}
	
	if forms, err = self.Parse(in, nil, &pos); err != nil {
		return err
	}

	return action(forms)
}

func (self *Scope) Let(key string, pos Pos) (int, error) {
	index := len(self.registers)

	if found := self.Get(key); found == nil {
		self.Set(key, Undefined)
		self.registers[key] = index
	} else if found.val != Undefined {
	        return index, Error(pos, "Attempt to override compile time binding: %v", key)
	} else if found.scope != self {
		found.Init(self, Undefined)
		self.registers[key] = index
	} else {
	        return index, Error(pos, "Duplicate binding: %v", key) 
	}

	return index, nil
}

func (self *Scope) Load(filePath string, stack *Stack) error {
	return self.Include(filePath, func(forms []Form) error {
		var ops []Op
		var err error
		
		if ops, err = NewScope(self).Compile(forms, nil); err != nil {
			return err
		}

		var registers Registers
		
		if err = EvalOps(ops, nil, registers[:], stack); err != nil {
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
		v, err := source.Get(n, pos)

		if err != nil {
			return err
		}

		if found := self.Get(n); found != nil {
			if v.dataType == &TFunction && found.val.dataType == &TFunction {
				v.data.(*Function).AddMethod(found.val.data.(*Function).methods...)
			} else if !useAll {
				return Error(pos, "Duplicate identifier: %v", n)
			}
		}

		self.Set(n, v)
	}

	return nil
}
