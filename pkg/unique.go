package gfoo

import (
	"fmt"
	"sync/atomic"
)

var nextUnique uint64

func Unique(prefix string) string {
	var suffix uint64
	
	for {
		suffix = nextUnique
		
		if atomic.CompareAndSwapUint64(&nextUnique, suffix, suffix+1) {
			break
		}
	}
	
	return fmt.Sprintf("%v-%v", prefix, suffix)
}

func (self *Scope) Unique(in string) string {
	if b, ok := self.bindings[in]; ok && b.val.data != nil {
		return b.val.data.(string)
	}

	out := Unique(in[1:])
	self.Set(in, NewVal(&TString, out))
	return out
}
