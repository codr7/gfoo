package gfoo

type PngModule struct {
	Scope
}

func (self *PngModule) Init() *Scope {
	self.Scope.Init()
	return &self.Scope
}
