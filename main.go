package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gfoo/pkg"
	"log"
	"os"
	"strings"
)

func repl(g *gfoo.GFoo) {
	fmt.Printf("gfoo v%v.%v\n\n", gfoo.VERSION_MAJOR, gfoo.VERSION_MINOR)
	fmt.Print("Press return on empty line to evaluate.\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	var buffer bytes.Buffer
	
	for {
		fmt.Print("  ")
		
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			break
		}
		
		line := scanner.Text()

		if line == "" {
			var err error
			source := buffer.String()
			buffer.Reset()
			in := bufio.NewReader(strings.NewReader(source))
			pos := gfoo.NewPos("repl")
			var forms []gfoo.Form
			
			if forms, err = g.Parse(in, &pos, nil); err != nil {
				fmt.Println(err)
				continue
			}

			scope := g.RootScope()
			var ops []gfoo.Op
			
			if ops, err = g.Compile(forms, scope, nil); err != nil {
				fmt.Println(err)
				continue
			}

			if err = g.Evaluate(ops, scope); err != nil {
				fmt.Println(err)
				continue
			}

			if err := g.DumpStack(os.Stdout); err != nil {
				log.Fatal(err)
			}
		
			fmt.Print("\n")
		} else if _, err := fmt.Fprintf(&buffer, "%v\n", line); err != nil {
			log.Fatal(err)
		}
	}	
}

func main() {
	g := gfoo.New()
	repl(g)
}
