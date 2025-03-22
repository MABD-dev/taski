package renderer

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mabd-dev/tasks/internal/models"
)

func RenderTable(tasks []models.Task) {
	table := table.New(os.Stdout)
	table.SetHeaders("#", "Name", "Description", "Status", "Created At")

	for _, task := range tasks {
		table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.Status.ToString(), task.CreatedAt.Format(time.RFC1123))
	}
	table.Render()
}
