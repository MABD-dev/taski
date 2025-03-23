package repl

import (
	"fmt"
	"os"
	"os/exec"
)

func help() error {
	fmt.Println("Welcome to taski REPL")
	fmt.Println("These are all the available commands")

	resetColor := "\033[0m"
	greenColor := "\033[32m"
	for _, cmd := range getCommands() {
		fmt.Printf(" - %v%v%v: %v\n", greenColor, cmd.name, resetColor, cmd.description)
	}

	return nil
}

func exit() error {
	os.Exit(0)
	return nil
}

func clear() error {
	// this is for linux only
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
