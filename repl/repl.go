package repl

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/mabd-dev/taski/internal/data/db"
)

func StartRepl() {
	fmt.Println("Taski REPL v0.1")

	db := db.GetDb()
	session := session{
		db: db,
	}

	scanner := bufio.NewScanner(os.Stdin)

	resetColor := "\033[0m"
	greenColor := "\033[32m"

	for {
		fmt.Printf("%vtaski>%v ", greenColor, resetColor)

		scanner.Scan()
		text := scanner.Text()
		input := splitInput(text)

		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		cmd := findCommand(commandName)
		if cmd == nil {
			fmt.Println("Invalid command")
			continue
		}

		cmd.handler(session, text)
	}
}

// split input by whitespace
func splitInput(text string) []string {
	loweredInput := strings.ToLower(text)
	return strings.Fields(loweredInput)
}

func findCommand(name string) *command {
	lowerName := strings.ToLower(name)

	for _, cmd := range getSortedCommands() {
		if cmd.name == lowerName || slices.Contains(cmd.alternativeNames, name) {
			return &cmd
		}
	}
	return nil
}
