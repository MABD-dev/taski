package repl

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mabd-dev/taski/internal/ui"
	"github.com/mabd-dev/taski/internal/domain/converter"
)

// TODO: this code is redundent with cmd package, find a way to combine them
func list(s session, input string) error {
	tasks := s.tasksRepo.List()

	// parsing status filtering flags
	statusValues := []string{}
	statusFunc := func(s string) error {
		statusValues = append(statusValues, s)
		return nil
	}

	listFlags := flag.NewFlagSet("list", flag.ContinueOnError)
	listFlags.Func("s", "Filter by status", statusFunc)

	parts := strings.Fields(input)
	listFlags.Parse(parts[1:])

	if len(statusValues) > 0 {
		statuses, err := converter.StringArrayToTaskStatus(statusValues)
		if err != nil {
			return err
		}
		tasks = converter.FilterByStatus(tasks, statuses)
	}

	if len(tasks) == 0 {
		if len(statusValues) > 0 {
			fmt.Println("No tasks smatch this filter")
		} else {
			fmt.Println("You don't have any tasks yet")
		}
		return nil
	} else {
		ui.RenderTable(tasks)
	}

	return nil
}

func help(s session, input string) error {
	fmt.Println("Welcome to taski REPL")
	fmt.Println("These are all the available commands")

	resetColor := "\033[0m"
	greenColor := "\033[32m"

	rawData := [][]string{}
	rawData = append(rawData, []string{"Command", "Description", "Alternatie Names"})

	for _, cmd := range getSortedCommands() {
		coloredName := greenColor + cmd.name + resetColor
		alternativeNames := strings.Join(cmd.alternativeNames, ", ")
		rawData = append(rawData, []string{coloredName, cmd.description, alternativeNames})
	}
	ui.RenderRawData(rawData)

	return nil
}

func exit(s session, input string) error {
	os.Exit(0)
	return nil
}

func clear(s session, input string) error {
	// this is for linux only
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
