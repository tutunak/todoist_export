package export

import (
	"strings"
	"testing"

	"github.com/tutunak/todoist_export/model"
)

func TestToMarkdown(t *testing.T) {
	data := &model.ExportData{
		Projects: []model.ProjectExport{
			{
				Name: "Project A",
				Tasks: []model.TaskExport{
					{
						Content:     "Task 1",
						IsCompleted: false,
						Priority:    4,
						Labels:      []string{"urgent"},
					},
					{
						Content:     "Task 2",
						IsCompleted: true,
						DueDate:     "2023-10-27",
					},
				},
			},
			{
				Name:  "Project B",
				Tasks: []model.TaskExport{},
			},
		},
	}

	md := ToMarkdown(data)

	// Check Project headers
	if !strings.Contains(md, "# Project A") {
		t.Error("Missing Project A header")
	}
	if !strings.Contains(md, "# Project B") {
		t.Error("Missing Project B header")
	}

	// Check Tasks
	if !strings.Contains(md, "- [ ] Task 1") {
		t.Error("Missing Task 1")
	}
	if !strings.Contains(md, "- [x] Task 2") {
		t.Error("Missing Task 2 (completed)")
	}

	// Check Metadata
	if !strings.Contains(md, "P4") {
		t.Error("Missing Priority P4")
	}
	if !strings.Contains(md, "Labels: urgent") {
		t.Error("Missing Label urgent")
	}
	if !strings.Contains(md, "Due: 2023-10-27") {
		t.Error("Missing Due Date")
	}

	// Check Empty Project
	if !strings.Contains(md, "_No tasks_") {
		t.Error("Missing _No tasks_ indicator for Project B")
	}
}
