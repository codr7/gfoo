package gfoo

import (
	"sync"
)

type Thread struct {
	body []Op
	stack, result Slice
	scope Scope
	done bool
	err error
	mutex sync.Mutex
	resume *sync.Cond
	resumeFlag bool
}

func NewThread(body []Op, scope *Scope) *Thread {
	t := new(Thread)
	t.body = body
	t.scope.Init(t)
	scope.Copy(&t.scope)
	t.resume = sync.NewCond(&t.mutex)
	return t
}

func (self *Thread) Call(stack *Slice, pos Pos) error {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.err != nil {
		return self.err
	}

	stack.Push(self.result.items...)
	self.result.Clear()
	
	if self.done {
		self.err = NewError(pos, "Thread is done")
	} else {
		self.resumeFlag = true
		self.resume.Signal()
	}
	
	return nil
}

func (self *Thread) Pause() {
	for !self.resumeFlag {
		self.resume.Wait()
	}

	self.resumeFlag = false
}

func (self *Thread) Start() {
	go func() {
		self.mutex.Lock()
		self.err = self.scope.Evaluate(self.body, &self.stack)
		self.result.items = self.stack.items
		self.done = true
		self.mutex.Unlock()
	}()
}
