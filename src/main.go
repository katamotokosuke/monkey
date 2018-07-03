package main

import (
	"os/user"
	"fmt"
	"repl"
	"os"
)

func main(){
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

