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

func repl(g *gfoo.Core, stack *gfoo.Slice) {
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

	flag.BoolVar(&g.Debug, "debug", false, "Enable debug mode")
	flag.Parse()

	args := flag.Args()
	argVals := gfoo.NewSlice(nil)
	g.Io.Set("ARGS", gfoo.NewVal(&gfoo.TSlice, argVals))
	
	stack := gfoo.NewSlice(nil)
	
	if len(args) == 0 {
		repl(g, stack)
	} else {
		for i := 1; i < len(args); i++ {
			argVals.Push(gfoo.NewVal(&gfoo.TString, args[i]))
		}

		if err := g.Load(args[0], stack); err != nil {
			log.Fatal(err)
		}
	}

	if err := g.Io.OUT.Flush(); err != nil {
		log.Fatal(err)
	}
}
