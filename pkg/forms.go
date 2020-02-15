package gfoo

var NilForms Forms

func init() {
	NilForms.Init(nil)
}

type Forms struct {
	items []Form
}

func (self *Forms) Init(items []Form) {
	self.items = items

	for i, j := 0, len(self.items)-1; i < j; i, j = i+1, j-1 {
		self.items[i], self.items[j] = self.items[j], self.items[i]
	}
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
