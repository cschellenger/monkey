package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/compiler"
	"github.com/cschellenger/monkey/evaluator"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"
	"github.com/cschellenger/monkey/vm"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

var input = `
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
fibonacci(35);
`

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	is := antlr.NewInputStream(input)
	l := parser.NewMonkeyLexer(is)
	stream := antlr.NewCommonTokenStream(l, 0)
	program := parser.NewMonkeyParser(stream)

	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program.Prog())
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())

		start := time.Now()

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()
	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program.Prog(), env)
		duration = time.Since(start)
	}

	fmt.Printf(
		"engine=%s, result=%s, duration=%s\n",
		*engine,
		result.Inspect(),
		duration)
}
