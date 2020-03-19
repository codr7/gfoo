package gfoo

type Module struct {
	Scope
}

func (self *Module) Init() {
	self.Scope.Init()
}

