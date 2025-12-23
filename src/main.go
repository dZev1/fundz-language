package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dZev1/fundz-language/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Fundz version SNAPSHOT01w26 %s . %s . %s\n", user.Name, user.Gid, user.HomeDir)
	fmt.Printf("Type '%s', '%s' '%s' or '%s' for more information\n", "help", "license", "credits", "copyright")

	repl.Start(os.Stdin, os.Stdout)
}