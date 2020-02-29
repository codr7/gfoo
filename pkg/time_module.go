package gfoo

import (
	"time"
)

func nowImp(stack *Slice, scope *Scope, pos Pos) (error) {
	stack.Push(NewVal(&TTime, time.Now().UTC()))
	return nil
}

func todayImp(stack *Slice, scope *Scope, pos Pos) (error) {
	now := time.Now().UTC()
	stack.Push(NewVal(&TTime, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)))
	return nil
}

func (self *Scope) InitTime() {
	s := new(Scope).Init()
	self.AddVal("time", &TScope, s)	

	s.AddVal("MAX", &TTime, MaxTime)
	s.AddVal("MIN", &TTime, MinTime)

	self.AddMethod("now", nil, []Result{RType(&TTime)}, nowImp)
	self.AddMethod("today", nil, []Result{RType(&TTime)}, todayImp)
}
