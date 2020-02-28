package gfoo

type Tree struct {
	compare TreeCompare
	root *TreeNode
	len uint64
}

type TreeNode struct {
	left, right *TreeNode
	key interface{}
	values []interface{}
	red bool
}

type TreeCompare = func(x, y interface{}) Order
type TreeNodeCompare = func(x, y *TreeNode) Order
type TreeAction = func(key, val interface{}) error

func NewTree(compare TreeCompare) *Tree {
	return new(Tree).Init(compare)
}

func (self *Tree) Init(compare TreeCompare) *Tree {
	self.compare = compare
	return self
}

func (self *Tree) Compare(other *Tree, compare TreeNodeCompare) Order {
	return self.root.Compare(other.root, compare)
}

func (self *Tree) Do(action TreeAction) error {
	return self.root.Do(action)
}

func (self *Tree) Find(key interface{}) []interface{} {
	n := self.findNode(self.root, key)
	if n == nil { return nil }
	return n.values
}

func (self *Tree) Insert(key, value interface{}, dup bool) bool {
	var ok bool
	self.root, ok = self.insertNode(self.root, key, value, dup)
	self.root.red = false
	return ok
}

func (self *Tree) Len() uint64 {
	return self.len
}

func (self *Tree) Update(key, value interface{}) bool {
	var ok bool
	self.root, ok = self.updateNode(self.root, key, value)
	self.root.red = false
	return ok
}

func (self *TreeNode) Compare(other *TreeNode, compare TreeNodeCompare) Order {
	if self == nil && other == nil {
		return Eq
	}
	
	if self == nil {
		return Lt
	}

	if other == nil {
		return Gt
	}

	if out := self.left.Compare(other.left, compare); out != Eq {
		return out
	}

	if out := compare(self, other); out != Eq {
		return out
	}

	if out := self.right.Compare(other.right, compare); out != Eq {
		return out
	}
	
	return Eq
}

func (self *TreeNode) Do(action TreeAction) (error) {
	if self != nil {
		if err := self.left.Do(action); err != nil {
			return err
		}
		
		for _, v := range self.values {
			if err := action(self.key, v); err != nil {
				return err
			}
		}
		
		if err := self.right.Do(action); err != nil {
			return err
		}
	}
	
	return nil
}

func (self *Tree) findNode(node *TreeNode, key interface{}) *TreeNode {
	for node != nil {
		switch self.compare(key, node.key) {
		case Lt:
			node = node.left
		case Gt:
			node = node.right
		default:
			return node
		}
	}

	return nil
}

func (self *Tree) insertNode(node *TreeNode, key, value interface{}, dup bool) (*TreeNode, bool) {
	if node == nil {
		node = &TreeNode{key: key, values: []interface{}{value}, red: true}
		self.len++
		return node, true
	}

	var ok bool

	switch self.compare(key, node.key) {
	case Lt:
		node.left, ok = self.insertNode(node.left, key, value, dup)
	case Gt:
		node.right, ok = self.insertNode(node.right, key, value, dup)
	default:
		if !dup {
			return node, false
		}
		
		node.values = append(node.values, value)
		self.len++
		return node, true
	}

	return node.fix(), ok
}

func (self *Tree) updateNode(node *TreeNode, key, value interface{}) (*TreeNode, bool) {
	if node == nil {
		node = &TreeNode{key: key, values: []interface{}{value}, red: true}
		self.len++
		return node, true
	}

	var ok bool

	switch self.compare(key, node.key) {
	case Lt:
		node.left, ok = self.updateNode(node.left, key, value)
	case Gt:
		node.right, ok = self.updateNode(node.right, key, value)
	default:
		if len(node.values) == 0 {
			node.values = append(node.values, value)
			self.len++
		} else {
			node.values[0] = value
		}
		
		return node, true
	}

	return node.fix(), ok
}

func (self *TreeNode) fix() *TreeNode {
	if (self.right.isRed()) {
		self = self.rotl()
	}

	if (self.left.isRed() && self.left.left.isRed()) {
		self = self.rotr()
	}

	if (self.left.isRed() && self.right.isRed()) {
		self.flip()
	}

	return self
}

func (self *TreeNode) flip() {
	self.red = !self.red
	self.left.red = !self.left.red
	self.right.red = !self.right.red
}

func (self *TreeNode) isRed() bool {
	return self != nil && self.red
}

func (self *TreeNode) rotl() *TreeNode {
	r := self.right
	self.right = r.left
	r.left = self
	r.red = self.red
	self.red = true
	return r
}

func (self *TreeNode) rotr() *TreeNode {
	l := self.left
	self.left = l.right
	l.right = self
	l.red = self.red
	self.red = true
	return l
}
