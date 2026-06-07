package repl

import (
	"io"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/evaluator"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"

	"github.com/chzyer/readline"
)

const PROMPT = ">> "
const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

var completer = readline.NewPrefixCompleter(
	readline.PcItem("puts"),
	readline.PcItem("len"),
	readline.PcItem("first"),
	readline.PcItem("last"),
	readline.PcItem("rest"),
	readline.PcItem("push"),
	readline.PcItem("fn"),
	readline.PcItem("while"),
)

func Start(in io.Reader, out io.Writer) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/monkey.readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	rl.CaptureExitSignal()

	env := object.NewEnvironment()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		input := antlr.NewInputStream(line)
		l := parser.NewMonkeyLexer(input)
		stream := antlr.NewCommonTokenStream(l, 0)
		p := parser.NewMonkeyParser(stream)
		errListener := evaluator.NewErrListener()
		p.AddErrorListener(errListener)

		program := p.Prog()
		if len(errListener.Errors) != 0 {
			printParserErrors(out, errListener.Errors)
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
