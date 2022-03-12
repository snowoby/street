package main

import (
	"fmt"
	"os"
	"street/cmd/makemigrations"
	"street/cmd/migrate"
	"street/cmd/street"
	"street/cmd/tasker"
)

type run = func()

func main() {
	argsWithProg := os.Args
	commandMap := map[string]run{
		"street":         street.Main,
		"tasker":         tasker.Main,
		"migrate":        migrate.Main,
		"makemigrations": makemigrations.Main,
	}
	command := "street"
	if len(argsWithProg) >= 2 {
		command = argsWithProg[1]
	}
	if run, ok := commandMap[command]; ok {
		run()
	} else {
		fmt.Printf("%s is not a valid command\n", command)
	}
}
