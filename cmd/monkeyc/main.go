package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/common"
	"github.com/cschellenger/monkey/compiler"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"
)

func main() {
	out := flag.String("o", "a.mvm", "outfile")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "%s [-o outfile] <infile>\n", os.Args[0])
		os.Exit(1)
	}
	outfile := *out

	infile := flag.Arg(0)
	file, err := os.Open(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %s: %v\n", infile, err)
		os.Exit(1)
	}
	defer func() {
		_ = file.Close()
	}()
	constants := []object.Object{}
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}
	input := antlr.NewIoStream(file)
	l := parser.NewMonkeyLexer(input)
	stream := antlr.NewCommonTokenStream(l, 0)
	p := parser.NewMonkeyParser(stream)
	errListener := common.NewErrListener()
	p.AddErrorListener(errListener)

	program := p.Prog()
	if len(errListener.Errors) != 0 {
		printParserErrors(errListener.Errors)
	}
	comp := compiler.NewWithState(symbolTable, constants)
	err = comp.Compile(program)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Woops! Compilation failed:\n %s\n", err)
	}
	toSave, err := os.Create(outfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %s for writing: %v\n", outfile, err)
	}
	defer func() { _ = toSave.Close() }()
	common.RegisterGobTypes()
	enc := gob.NewEncoder(toSave)
	if err = enc.Encode(constants); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to encode constants: %v\n", err)
		os.Exit(1)
	}
	if err = enc.Encode(comp.Bytecode()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to encode bytecode: %v\n", err)
		os.Exit(1)
	}
}

func printParserErrors(errors []string) {
	io.WriteString(os.Stderr, common.MONKEY_FACE)
	io.WriteString(os.Stderr, "Woops! We ran into some monkey business here!\n")
	io.WriteString(os.Stderr, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(os.Stderr, "\t"+msg+"\n")
	}
}
