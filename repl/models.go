package repl

import "github.com/mabd-dev/taski/internal/data/db"

type session struct {
	db db.Db
}

type command struct {
	name             string
	alternativeNames []string
	description      string
	handler          func(s session, input string) error
}

func getSortedCommands() []command {
	return []command{
		{
			name:             "list",
			alternativeNames: []string{"ls"},
			description:      "List your tasks",
			handler:          list,
		},
		{
			name:             "clear",
			alternativeNames: []string{"cls"},
			description:      "Clear the terminal",
			handler:          clear,
		},
		{
			name:             "exit",
			alternativeNames: []string{"kill"},
			description:      "Stops the program",
			handler:          exit,
		},
		{
			name:             "help",
			alternativeNames: []string{"h"},
			description:      "Show list of available commands",
			handler:          help,
		},
	}
}
