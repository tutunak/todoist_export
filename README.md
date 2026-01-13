# Todoist Export

A CLI tool for exporting your Todoist projects and tasks into a structured YAML format.

## Features
- **Project grouping**: Exports tasks grouped by their respective projects (including Inbox).
- **Label Resolution**: Resolves label IDs to their human-readable names.
- **Rich Data**: Includes task content, description, priority, due dates, and completion status.
- **YAML Output**: Produces clean, readable YAML.

## Prerequisites
- **Go 1.22+**
- **Todoist API Token**: You can get this from the [Todoist Integrations Settings](https://todoist.com/prefs/integrations).

## Installation

### From Source
1. Clone the repository:
   ```bash
   git clone https://github.com/tutunak/todoist_export.git
   cd todoist_export
   ```
2. Build the binary:
   ```bash
   go build -o todoist-export main.go
   ```

## Usage

The application requires the `TODOIST_API_TOKEN` environment variable to be set.

### Running with Go
```bash
export TODOIST_API_TOKEN=your_token_here

# Default behavior (exports to <timestamp>.yaml)
go run .

# Export to Markdown
go run . -format md

# Export to specific file
go run . -output my_backup.yaml
```

### Running the Binary
```bash
export TODOIST_API_TOKEN=your_token_here

# Default (exports to <timestamp>.yaml)
./todoist-export

# Custom options
./todoist-export -format md -output todoist_backup.md
```

## Output Format
The output is a YAML document with the following structure:

```yaml
projects:
  - id: "2255441111"
    name: "Inbox"
    tasks:
      - content: "Buy milk"
        description: "2% fat"
        priority: 1
        labels:
          - "groceries"
        due_date: "2023-10-27"
        url: "https://todoist.com/showTask?id=..."
        is_completed: false
```

## Testing
Run the unit tests:
```bash
go test ./...
```