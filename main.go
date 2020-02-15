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

func repl(g *gfoo.Scope) {
	fmt.Printf("gfoo v%v.%v\n\n", gfoo.VersionMajor, gfoo.VersionMinor)
	fmt.Print("Press Return on empty line to evaluate.\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	var buffer bytes.Buffer
	stack := gfoo.NewSlice(nil)
	
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

			var ops []gfoo.Op
			
			if ops, err = g.Compile(forms, nil); err != nil {
				fmt.Println(err)
				continue
			}

			if err = g.Evaluate(ops, stack); err != nil {
				fmt.Println(err)
				continue
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
	g := gfoo.New()
	repl(g)
}
