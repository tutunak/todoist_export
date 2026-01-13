package model



// Todoist Entities

type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IsFavorite bool   `json:"is_favorite"`
	ViewStyle  string `json:"view_style"`
}

type Task struct {
	ID           string   `json:"id"`
	ProjectID    string   `json:"project_id"`
	SectionID    string   `json:"section_id"`
	Content      string   `json:"content"`
	Description  string   `json:"description"`
	IsCompleted  bool     `json:"is_completed"`
	Labels       []string `json:"labels"`
	Priority     int      `json:"priority"`
	CommentCount int      `json:"comment_count"`
	CreatorID    string   `json:"creator_id"`
	CreatedAt    string   `json:"created_at"`
	Due          *Due     `json:"due"`
	URL          string   `json:"url"`
}

type Due struct {
	Date      string `json:"date"`
	String    string `json:"string"`
	Lang      string `json:"lang"`
	IsRecurring bool `json:"is_recurring"`
}

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Export Entities

type ExportData struct {
	Projects []ProjectExport `yaml:"projects"`
}

type ProjectExport struct {
	ID    string       `yaml:"id"`
	Name  string       `yaml:"name"`
	Tasks []TaskExport `yaml:"tasks"`
}

type TaskExport struct {
	Content     string   `yaml:"content"`
	Description string   `yaml:"description,omitempty"`
	Priority    int      `yaml:"priority,omitempty"`
	Labels      []string `yaml:"labels,omitempty"` // Mapped names
	DueDate     string   `yaml:"due_date,omitempty"`
	URL         string   `yaml:"url,omitempty"`
    IsCompleted bool     `yaml:"is_completed"`
}
