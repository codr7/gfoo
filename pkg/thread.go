package gfoo

import (
	"sync"
)

type Thread struct {
	body []Op
	stack Slice
	scope Scope
	done bool
	err error
	mutex sync.Mutex
}

func NewThread(body []Op, scope *Scope) *Thread {
	t := new(Thread)
	t.body = body
	t.scope.Init(scope.vm, scope.thread)
	scope.Copy(&t.scope, false)
	return t
}

func (self *Thread) Start() {
	go func() {
		self.mutex.Lock()
		self.err = self.scope.Evaluate(self.body, &self.stack)
		self.done = true
		self.mutex.Unlock()
	}()
}

func (self *Thread) Call(stack *Slice) error {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	
	if self.err != nil {
		return self.err
	}

	if self.done {
		stack.Push(self.stack.items...)
	}
	
	return nil
}
