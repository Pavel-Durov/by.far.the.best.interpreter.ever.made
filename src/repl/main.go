package repl

import (
	"bufio"
	"fmt"
	"io"

	"by.far.the.best.interpreter.ever.made.io/src/lexer"
)

type TokenType string

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		for tok := lex.NextToken(); !tok.IsEOF(); tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
