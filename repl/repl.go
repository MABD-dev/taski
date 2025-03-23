package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	fmt.Println("Taski REPL v0.1")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()
		input := splitInput(text)

		if len(input) == 0 {
			continue
		}

		command := input[0]

		cmd, ok := getCommands()[command]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		cmd.handler()
	}
}

// split input by whitespace
func splitInput(text string) []string {
	loweredInput := strings.ToLower(text)
	return strings.Fields(loweredInput)
}
