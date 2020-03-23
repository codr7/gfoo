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

func (self *Scope) Unique(key string) string {
	if b := self.Get(key); b != nil && b.val.data != nil {
		return b.val.data.(string)
	}

	out := Unique(key[1:])
	self.Set(key, NewVal(&TString, out))
	return out
}
