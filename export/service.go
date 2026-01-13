package export

import (
	"fmt"


	"github.com/tutunak/todoist_export/model"
	"github.com/tutunak/todoist_export/todoist"
)

type Service struct {
	client *todoist.Client
}

func NewService(client *todoist.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Export() (*model.ExportData, error) {
	// 1. Fetch Projects
	projects, err := s.client.GetProjects()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch projects: %w", err)
	}

	// 2. Fetch all Tasks
	allTasks, err := s.client.GetTasks("")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	// 3. Fetch Labels (to map IDs to names)
	labels, err := s.client.GetLabels()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch labels: %w", err)
	}
	labelMap := make(map[string]string)
	for _, l := range labels {
		labelMap[l.ID] = l.Name
	}

	// Group tasks by Project ID
	tasksByProject := make(map[string][]model.Task)
	for _, t := range allTasks {
		tasksByProject[t.ProjectID] = append(tasksByProject[t.ProjectID], t)
	}

	// Build Project Exports
	var projectExports []model.ProjectExport
	
	// Track processed projects to handle "Inbox" explicitly if needed, but Todoist Inbox is just a project.
	// We want to ensure all projects found are exported.
	
	for _, p := range projects {
		tasks := tasksByProject[p.ID]
		// Convert model.Task to model.TaskExport
		exportTasks := convertTasks(tasks, labelMap)
		
		projectExports = append(projectExports, model.ProjectExport{
			ID:    p.ID,
			Name:  p.Name,
			Tasks: exportTasks,
		})
	}

	return &model.ExportData{
		Projects: projectExports,
	}, nil
}

func convertTasks(tasks []model.Task, labelMap map[string]string) []model.TaskExport {
	var exports []model.TaskExport
	for _, t := range tasks {
		var labels []string
		for _, lID := range t.Labels {
			if name, ok := labelMap[lID]; ok {
				labels = append(labels, name)
			} else {
				labels = append(labels, lID) // Fallback to ID
			}
		}

		dueDate := ""
		if t.Due != nil {
			dueDate = t.Due.Date // YYYY-MM-DD or RFC3339
		}

		exports = append(exports, model.TaskExport{
			Content:     t.Content,
			Description: t.Description,
			Priority:    t.Priority,
			Labels:      labels,
			DueDate:     dueDate,
			URL:         t.URL,
			IsCompleted: t.IsCompleted,
		})
	}
	return exports
}
