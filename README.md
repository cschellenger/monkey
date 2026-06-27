# Monkey Programming Language

## Description
This is my implementation of the books ["Writing An Interpreter In Go"](https://interpreterbook.com/) and ["Writing a Compiler in Go"](https://compilerbook.com/) by Thorsten Ball. A large amount of code is taken directly from the books, but I have made the following additions.
- Support for `float` data type
- Support for `while` loops
- Operator overloading - not yet user feature because defining new classes / objects is not supported
- GNU readline support in REPL
- Replaced hand-written lexer and parser with antlr based lexer-parser

## Core Language TODOs
- Support for-in loops
- Support user defined classes / objects

## Compiler Progress
- Finished with minimal compiler and VM
- Ready for Compiling Expressions chapter

## Requirements for Building
The following tools must be in your `$PATH` to build.

- go
- java - only required for running antlr

