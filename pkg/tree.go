package gfoo

type Tree struct {
	compare TreeKeyCompare
	root *TreeNode
	len uint64
}

type TreeNode struct {
	left, right *TreeNode
	key interface{}
	values []interface{}
	red bool
}

type TreeAction = func(key, val interface{}) error
type TreeKeyCompare = func(x, y interface{}) Order
type TreeNodeCompare = func(x, y *TreeNode) Order

func NewTree(compare TreeKeyCompare) *Tree {
	return new(Tree).Init(compare)
}

func (self *Tree) Init(compare TreeKeyCompare) *Tree {
	self.compare = compare
	return self
}

func (self *Tree) Compare(other *Tree, compare TreeNodeCompare) Order {
	return self.root.Compare(other.root, compare)
}

func (self *Tree) Find(key interface{}) []interface{} {
	n := self.findNode(self.root, key)
	if n == nil { return nil }
	return n.values
}

func (self *Tree) ForEach(action TreeAction) error {
	return self.root.ForEach(action)
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

func (self Tree) Union(source *Tree) Tree {
	self.root = self.unionNode(self.root.clone(), source.root)
	self.root.red = false
	return self
}

func (self Tree) Update(key, value interface{}) Tree {
	self.root = self.updateNode(self.root.clone(), key, value)
	self.root.red = false
	return self
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

func (self *TreeNode) ForEach(action TreeAction) (error) {
	if self != nil {
		if err := self.left.ForEach(action); err != nil {
			return err
		}
		
		for _, v := range self.values {
			if err := action(self.key, v); err != nil {
				return err
			}
		}
		
		if err := self.right.ForEach(action); err != nil {
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

func (self *Tree) unionNode(target, source *TreeNode) *TreeNode {
	if source == nil {
		return target
	}
	
	self.root = self.unionNode(self.root, source.left)
	self.root = self.unionNode(self.root, source.right)

	if target == nil {
		target = &TreeNode{
			key: source.key,
			values: make([]interface{}, len(source.values)),
			red: true}
		copy(target.values, source.values)
		self.len += uint64(len(source.values))
		self.root = self.unionNode(self.root, source.left)
		self.root = self.unionNode(self.root, source.right)
		return target
	}

	switch self.compare(source.key, target.key) {
	case Lt:
		target.left = self.unionNode(target.left.clone(), source)
	case Gt:
		target.right = self.unionNode(target.right.clone(), source)
	}

	return target.fix()
}

func (self *Tree) updateNode(node *TreeNode, key, value interface{}) *TreeNode {
	if node == nil {
		node = &TreeNode{key: key, values: []interface{}{value}, red: true}
		self.len++
		return node
	}

	switch self.compare(key, node.key) {
	case Lt:
		node.left = self.updateNode(node.left.clone(), key, value)
	case Gt:
		node.right = self.updateNode(node.right.clone(), key, value)
	default:
		if len(node.values) == 0 {
			node.values = append(node.values, value)
			self.len++
		} else {
			node.values[0] = value
		}
		
		return node
	}

	return node.fix()
}

func (self *TreeNode) clone() *TreeNode {
	if self == nil {
		return nil
	}
	
	out := new(TreeNode)
	out.left = self.left
	out.right = self.right
	out.key = self.key
	out.red = self.red

	if len(self.values) > 0 {
		out.values = make([]interface{}, len(self.values))
		copy(out.values, self.values)	
	}

	return out
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
