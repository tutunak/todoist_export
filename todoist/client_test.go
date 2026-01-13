package todoist

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjects(t *testing.T) {
	mockResponse := `[
		{"id": "123", "name": "Inbox", "is_favorite": false},
		{"id": "456", "name": "Work", "is_favorite": true}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/projects" {
			t.Errorf("Expected path /projects, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test_token" {
			t.Errorf("Expected Authorization header, got %s", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewClient("test_token")
	client.BaseURL = server.URL

	projects, err := client.GetProjects()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(projects) != 2 {
		t.Errorf("Expected 2 projects, got %d", len(projects))
	}
	if projects[0].Name != "Inbox" {
		t.Errorf("Expected first project name Inbox, got %s", projects[0].Name)
	}
}
