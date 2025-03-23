package repl

type CommandHandler interface {
	Run() error
}

type command struct {
	name        string
	description string
	handler     func() error
}

func getCommands() map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "Stops the program",
			handler:     exit,
		},
		"help": {
			name:        "help",
			description: "Show list of available commands",
			handler:     help,
		},
		"clear": {
			name:        "clear",
			description: "clear the terminal",
			handler:     clear,
		},
	}
}
