package main

import (
	"fmt"
	"gorila/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Gorila programming language!\n", usr.Username)

	fmt.Printf("Fell fere to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)

}
