package gfoo

type Type interface {
	DirectParents() []Type
	Name() string
}

type TypeBase struct {
 	directParents []Type
	name string
	parents map[Type]Type
}

func NewType(name string, parents...Type) Type {
	return new(TypeBase).Init(name, parents)
}

func (self *TypeBase) Init(name string, parents []Type) *TypeBase {
	self.name = name
	self.parents = make(map[Type]Type)
	
	for _, p := range parents {
		self.Derive(p)
	}
	
	return self
}

func (self *TypeBase) derive(parent, direct Type) {
	for _, p := range parent.DirectParents() {
		self.derive(p, direct)
	}

	self.parents[parent] = direct
}

func (self *TypeBase) Derive(parent Type) {
	self.derive(parent, parent)
	self.directParents = append(self.directParents, parent)
}

func (self *TypeBase) DirectParents() []Type {
	return self.directParents
}


func (self *TypeBase) Name() string {
	return self.name
}
