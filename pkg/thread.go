package gfoo

type Thread struct {
	body []Op
	err error
	results chan []Val
	stack Stack
}

func NewThread(body []Op) *Thread {
	t := new(Thread)
	t.body = body
	t.results = make(chan []Val, 0) 
	return t
}

func (self *Thread) Pause(result []Val) {
	self.results<- result
}

func (self *Thread) Start() {
	go func() {
		if self.err = EvalOps(self.body, self, NewStack(nil), &self.stack); self.err == nil {
			self.results<- self.stack.items
		}
		
		close(self.results)
	}()
}

func (self *Thread) Wait(stack *Stack, pos Pos) error {
	if result, ok := <-self.results; ok {
		stack.Push(result...)
	} else {
		if self.err != nil {
			return self.err
		}

		return Error(pos, "Thread is done")
	}
		
	return nil
}
