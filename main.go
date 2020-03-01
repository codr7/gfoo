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

func repl(g *gfoo.Scope, stack *gfoo.Slice) {
	fmt.Printf("gfoo v%v.%v\n\n", gfoo.VersionMajor, gfoo.VersionMinor)
	fmt.Print("Press Return on empty line to evaluate.\n\n")

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
			source := strings.TrimSpace(buffer.String());
			buffer.Reset()

			if source == "" {
				stack.Clear()
			} else {
				in := bufio.NewReader(strings.NewReader(source))
				pos := gfoo.NewPos("repl")
				var forms []gfoo.Form
			
				if forms, err = g.Parse(in, nil, &pos); err != nil {
					fmt.Println(err)
					continue
				}
				
				var ops []gfoo.Op
				
				if ops, err = g.Compile(forms, nil); err != nil {
					fmt.Println(err)
					continue
				}
				
				if err = g.EvalOps(ops, stack); err != nil {
					fmt.Println(err)
				}
			}
			
			if err := stack.Dump(os.Stdout); err != nil {
				log.Fatal(err)
			}
		
			fmt.Print("\n")
		} else if _, err := fmt.Fprintf(&buffer, "%v\n", line); err != nil {
			log.Fatal(err)
		}
	}	
}

func main() {
	gfoo.Init()
	g := gfoo.New()
	g.Debug = true
	stack := gfoo.NewSlice(nil)

	if len(os.Args) == 1 {
		repl(g, stack)
	} else {
		for _, path := range os.Args[1:] {
			if err := g.Load(path, stack); err != nil {
				log.Fatal(err)
			}
		}
	}	
}
