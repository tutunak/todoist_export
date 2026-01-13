package export

import (
	"testing"
	"net/http"
	"net/http/httptest"


	"github.com/tutunak/todoist_export/todoist"
)

func TestExport(t *testing.T) {
	// Mock Todoist API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/projects":
			w.Write([]byte(`[{"id": "p1", "name": "Project1"}, {"id": "p2", "name": "Project2"}]`))
		case "/tasks":
			// Check if filter param is present
			filter := r.URL.Query().Get("filter")
			if filter == "today" {
				w.Write([]byte(`[{"id": "t3", "content": "FilterTask", "project_id": "p1"}]`))
			} else {
				w.Write([]byte(`[
					{"id": "t1", "content": "Task1", "project_id": "p1", "labels": ["l1"]},
					{"id": "t2", "content": "Task2", "project_id": "p2"}
				]`))
			}
		case "/labels":
			w.Write([]byte(`[{"id": "l1", "name": "urgent"}]`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := todoist.NewClient("token")
	client.BaseURL = server.URL

	svc := NewService(client)
	data, err := svc.Export()
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	// Verify Projects
	if len(data.Projects) != 2 {
		t.Errorf("Expected 2 projects, got %d", len(data.Projects))
	}

	p1 := data.Projects[0] // Order might depend on map iteration if not sorted, but projects list from API is slice.
	// Wait, loops over projects slice, so order preserved from API response (if stable) or map lookup. 
	// Projects slice comes from GetProjects.
	if p1.ID == "p1" {
		if len(p1.Tasks) != 1 {
			t.Errorf("Expected 1 task in Project1, got %d", len(p1.Tasks))
		}
		if p1.Tasks[0].Content != "Task1" {
			t.Errorf("Expected Task1, got %s", p1.Tasks[0].Content)
		}
		if len(p1.Tasks[0].Labels) != 1 || p1.Tasks[0].Labels[0] != "urgent" {
			t.Errorf("Expected label 'urgent', got %v", p1.Tasks[0].Labels)
		}
	}
}
