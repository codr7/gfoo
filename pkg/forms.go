package gfoo

type Forms struct {
	items []Form
}

func (self *Forms) Init(items []Form) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}

	self.items = items
}

func (self *Forms) Len() int {
	return len(self.items)
}

func (self *Forms) Pop() Form {
	i := len(self.items)-1

	if i == -1 {
		return nil
	}
	
	f := self.items[i]
	self.items = self.items[:i]
	return f
}
