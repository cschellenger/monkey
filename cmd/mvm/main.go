package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"

	"github.com/cschellenger/monkey/common"
	"github.com/cschellenger/monkey/compiler"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/vm"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "must provide an input mvm file\n")
		os.Exit(1)
	}
	infile, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file %s: %v\n", flag.Arg(1), err)
		os.Exit(1)
	}
	common.RegisterGobTypes()
	enc := gob.NewDecoder(infile)

	constants := []object.Object{}
	bytecode := compiler.Bytecode{}
	if err = enc.Decode(&constants); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read constants: %v\n", err)
		os.Exit(1)
	}
	if err = enc.Decode(&bytecode); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read bytecode: %v\n", err)
		os.Exit(1)
	}
	_ = infile.Close()
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}
	machine := vm.NewWithGlobalsStore(&bytecode, globals)
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running monkey program: %v\n", err)
		os.Exit(1)
	}
}
