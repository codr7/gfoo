package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"gfoo/pkg"
	"log"
	"os"
	"strings"
)

func repl(g *gfoo.Core, stack *gfoo.Stack) {
	fmt.Printf("gfoo v%v.%v\n\n", gfoo.VersionMajor, gfoo.VersionMinor)
	fmt.Print("Press Return on empty line to evaluate.\n\n")
	
	scanner := bufio.NewScanner(os.Stdin)
	var buffer bytes.Buffer
	registers := gfoo.NewStack(nil)
	g.Eval("use: abc...", nil, registers, nil)
	
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
				if err = g.Eval(source, nil, registers, stack); err != nil {
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

	flag.BoolVar(&gfoo.Debug, "debug", false, "Enable debug mode")
	flag.Parse()

	args := flag.Args()
	stack := gfoo.NewStack(nil)
	
	if len(args) == 0 {
		repl(g, stack)
	} else {
		for i := 1; i < len(args); i++ {
			g.Io.ARGS.Push(gfoo.NewVal(&gfoo.TString, args[i]))
		}

		if err := g.Load(args[0], stack); err != nil {
			log.Fatal(err)
		}
	}

	if err := g.Io.OUT.Flush(); err != nil {
		log.Fatal(err)
	}
}
