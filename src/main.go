package main

import (
	"fmt"
	"os"
	"os/user"

	"by.far.the.best.interpreter.ever.made.io/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s\n", user.Username)
	fmt.Println("Welcome to the best REPL ever made!")
	repl.Start(os.Stdin, os.Stdout)
	fmt.Println("Hello, world!")
}
