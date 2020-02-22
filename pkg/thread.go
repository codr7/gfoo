package gfoo

type Thread struct {
	body []Op
	stack Slice
	scope Scope
	done bool
	err error
	results chan []Val
}

func NewThread(body []Op, scope *Scope) *Thread {
	t := new(Thread)
	t.body = body
	t.scope.Init()
	scope.Copy(&t.scope)
	t.scope.thread = t
	t.results = make(chan []Val, 0) 
	return t
}

func (self *Thread) Call(stack *Slice, pos Pos) error {
	if result, ok := <-self.results; ok {
		stack.Push(result...)
	} else {
		if self.err != nil {
			return self.err
		}

		stack.Push(self.stack.items...)
		self.err = NewError(pos, "Thread is done")
	}
		
	return nil
}

func (self *Thread) Pause(result []Val) {
	self.results<- result
}

func (self *Thread) Start() {
	go func() {	
		self.err = self.scope.Evaluate(self.body, &self.stack)
		close(self.results)
	}()
}
