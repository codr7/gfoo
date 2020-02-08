package gfoo

type Forms struct {
	items []Form
}

func (self *Forms) Init(items []Form) {
	self.items = items
}

func (self *Forms) Len() int {
	return len(self.items)
}

func (self *Forms) Pop() Form {
	if len(self.items) == 0 {
		return nil
	}
	
	f := self.items[0]
	self.items = self.items[1:]
	return f
}
