package export

import (
	"fmt"
	"strings"

	"github.com/tutunak/todoist_export/model"
)

// ToMarkdown converts ExportData to a Markdown string
func ToMarkdown(data *model.ExportData) string {
	var sb strings.Builder

	for _, project := range data.Projects {
		sb.WriteString(fmt.Sprintf("# %s\n\n", project.Name))

		if len(project.Tasks) == 0 {
			sb.WriteString("_No tasks_\n\n")
			continue
		}

		for _, task := range project.Tasks {
			checkbox := "- [ ]"
			if task.IsCompleted {
				checkbox = "- [x]"
			}

			// Title and Link
			line := fmt.Sprintf("%s %s", checkbox, task.Content)
			if task.URL != "" {
				line += fmt.Sprintf(" ([Link](%s))", task.URL)
			}
			sb.WriteString(line + "\n")

			// Metadata block
			var meta []string
			if task.DueDate != "" {
				meta = append(meta, fmt.Sprintf("Due: %s", task.DueDate))
			}
			if task.Priority > 1 {
				// Todoist priority 4 is highest (p1 in UI), 1 is lowest.
				// Let's just show it as number for now or map it if needed.
				// Todoist API: 4=p1, 1=p4.
				meta = append(meta, fmt.Sprintf("P%d", task.Priority))
			}
			if len(task.Labels) > 0 {
				meta = append(meta, fmt.Sprintf("Labels: %s", strings.Join(task.Labels, ", ")))
			}

			if len(meta) > 0 {
				sb.WriteString(fmt.Sprintf("  > %s\n", strings.Join(meta, " | ")))
			}

			// Description
			if task.Description != "" {
				// Indent description
				lines := strings.Split(task.Description, "\n")
				for _, l := range lines {
					sb.WriteString(fmt.Sprintf("  %s\n", l))
				}
			}
			sb.WriteString("\n")
		}
		sb.WriteString("---\n\n")
	}

	return sb.String()
}
