package main

import (
	"gfoo/pkg"
	"os"
)

func repl(gfoo *gfoo.GFoo) {
}

func main() {
	gfoo := gfoo.New()
	gfoo.Symbol("foo").Quote().Dump(os.Stdout)
	repl(gfoo)
}
