package gfoo

type Bindings = map[string]Binding

type Scope struct {
	vm *VM
	thread *Thread
	bindings Bindings
}

func (self *Scope) Init(vm *VM, thread *Thread) *Scope {
	self.vm = vm
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
	for k, b := range self.bindings {
		if !child && b.scope == self {
			b.scope = out
		}
		
		out.bindings[k] = b
	}
}

func (self *Scope) Clone(child bool) *Scope {
	out := new(Scope).Init(self.vm, self.thread)
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

func (self *Scope) Set(key string, val Val) {
	self.bindings[key] = NewBinding(self, val)
}
