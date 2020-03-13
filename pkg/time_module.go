package gfoo

import (
	"time"
)

func daysImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TTimeDelta, Days(int(stack.Pop().data.(*Int).Int64()))))
	return nil
}

func nowImp(stack *Slice, scope *Scope, pos Pos) error {
	stack.Push(NewVal(&TTime, time.Now().UTC()))
	return nil
}

func timeAddImp(stack *Slice, scope *Scope, pos Pos) error {
	d := stack.Pop().data.(TimeDelta)
	stack.Push(NewVal(&TTime, stack.Pop().data.(time.Time).AddDate(d.years, d.months, d.days)))
	return nil
}

func todayImp(stack *Slice, scope *Scope, pos Pos) error {
	now := time.Now().UTC()
	stack.Push(NewVal(&TTime, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)))
	return nil
}

func (self *Scope) InitTime() *Scope {
	self.AddType(&TTime)
	self.AddType(&TTimeDelta)

	self.AddVal("MAX", &TTime, MaxTime)
	self.AddVal("MIN", &TTime, MinTime)

	self.AddMethod("days", []Arg{AType("n", &TInt)}, []Ret{RType(&TTimeDelta)}, daysImp)
	self.AddMethod("now", nil, []Ret{RType(&TTime)}, nowImp)
	self.AddMethod("+", []Arg{AType("x", &TTime), AType("y", &TTimeDelta)}, []Ret{RType(&TTime)}, timeAddImp)
	self.AddMethod("today", nil, []Ret{RType(&TTime)}, todayImp)
	return self
}
