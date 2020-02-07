package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gfoo/pkg"
	"log"
	"os"
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
			var forms []gfoo.Form
			source := buffer.String()
			buffer.Reset()
			
			if forms, err = g.Parse(source); err != nil {
				fmt.Println(err)
				continue
			}

			var ops []gfoo.Op
			
			if ops, err = g.Compile(forms); err != nil {
				fmt.Println(err)
				continue
			}

			if err = g.Eval(ops); err != nil {
				fmt.Println(err)
				continue
			}

			if err := g.DumpStack(os.Stdout); err != nil {
				log.Fatal(err)
			}
		
			fmt.Print("\n")
		} else if _, err := buffer.WriteString(line); err != nil {
			log.Fatal(err)
		}
	}	
}

func main() {
	g := gfoo.New()
	repl(g)
}
