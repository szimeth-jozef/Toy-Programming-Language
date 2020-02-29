package main

import (
	"fmt"
	"os"
	"os/user"
	"Interpreter/shell"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the shell of the learning interpreter/language!\n", user.Username)
	shell.Start(os.Stdin, os.Stdout)
}
