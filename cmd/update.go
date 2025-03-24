package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strconv"

	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task name or description",
	Long:  "Update a task name or description by providing it's number",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// flags, name, description, status
		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		task := repos.TasksRepo.Get(taskNumber)
		if task == nil {
			return errors.New("could not find task!")
		}

		err = openTaskEditor(task)
		if err != nil {
			return err
		}

		err = repos.TasksRepo.Update(taskNumber, &task.Name, &task.Description, &task.Status)
		if err != nil {
			return err
		}

		ui.RenderKanbanBoard(repos.TasksRepo.GetAll())
		return nil
	},
}

func openTaskEditor(task *models.Task) error {
	tmpFile, err := os.CreateTemp("", "edit-task.json")
	if err != nil {
		return err
	}

	defer os.Remove(tmpFile.Name())

	jsonData, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		return err
	}

	if _, err := tmpFile.Write(jsonData); err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}

	// Determine the user's preferred text editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	// Open the text editor
	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	editedTextBytes, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return err
	}

	if err := json.Unmarshal(editedTextBytes, task); err != nil {
		return err
	}

	return nil
}
