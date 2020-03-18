package gfoo

type ImageModule struct {
	Scope
}

func newRgbaImp(scope *Scope, stack *Slice, pos Pos) error {
	h := stack.Pop().data.(Int)
	w := stack.Pop().data.(Int)
	stack.Push(NewVal(&TRgba, NewRgba(int(w), int(h))))
	return nil
}

func (self *ImageModule) Init() *Scope {
	self.Scope.Init()
	self.AddType(&TRgba)

	self.AddMethod("new-rgba",
		[]Arg{AType("width", &TInt), AType("height", &TInt)},
		[]Ret{RType(&TRgba)},
		newRgbaImp)
	
	return &self.Scope
}
