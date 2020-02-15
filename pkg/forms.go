package gfoo

type Forms struct {
	items []Form
}

func NewForms(items []Form) *Forms {
	return new(Forms).Init(items)
}

func (self *Forms) Init(items []Form) *Forms {
	self.items = items

	for i, j := 0, len(self.items)-1; i < j; i, j = i+1, j-1 {
		self.items[i], self.items[j] = self.items[j], self.items[i]
	}

	return self
}

func (self *Forms) Len() int {
	return len(self.items)
}

func (self *Forms) Pop() Form {
	i := len(self.items)
	
	if i == 0 {
		return nil
	}

	i--
	f := self.items[i]
	self.items = self.items[:i]
	return f
}

func (self *Forms) Push(form Form) {
	self.items = append(self.items, form)
}
