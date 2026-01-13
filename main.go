package main


import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/tutunak/todoist_export/export"

	"github.com/tutunak/todoist_export/todoist"
)

func main() {
	format := flag.String("format", "yaml", "Output format: yaml (default) or md/markdown")
	output := flag.String("output", "", "Output filename (default: auto-generated timestamp)")
	flag.Parse()

	token := os.Getenv("TODOIST_API_TOKEN")
	if token == "" {
		log.Fatal("TODOIST_API_TOKEN environment variable is required")
	}

	client := todoist.NewClient(token)
	svc := export.NewService(client)

	log.Println("Starting Todoist export...")
	data, err := svc.Export()
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}

	ext := "yaml"
	if *format == "md" || *format == "markdown" {
		ext = "md"
	}

	filename := *output
	if filename == "" {
		filename = fmt.Sprintf("%d.%s", time.Now().Unix(), ext)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filename, err)
	}
	defer f.Close()

	if ext == "md" {
		mdContent := export.ToMarkdown(data)
		if _, err := f.WriteString(mdContent); err != nil {
			log.Fatalf("Failed to write Markdown to file: %v", err)
		}
	} else {
		encoder := yaml.NewEncoder(f)
		encoder.SetIndent(2)
		if err := encoder.Encode(data); err != nil {
			log.Fatalf("Failed to encode YAML: %v", err)
		}
	}
	
	log.Printf("Export completed successfully to %s", filename)
}
