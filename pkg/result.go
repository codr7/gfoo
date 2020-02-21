package gfoo

type Result struct {
	Argument
}

func NewResult(id string, val Val) Result {
	var r Result
	return r.Init(id, val)
}

func (self Result) Init(id string, val Val) Result {
	self.Argument.Init(id, val)
	return self
}
