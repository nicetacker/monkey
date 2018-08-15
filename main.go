package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nicetacker/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey programming language!\n",
		user.Username)
	fmt.Printf("eel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
