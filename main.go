package main

import (

	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/tutunak/todoist_export/export"
	"github.com/tutunak/todoist_export/todoist"
)

func main() {
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

	encoder := yaml.NewEncoder(os.Stdout)
	encoder.SetIndent(2)
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Failed to encode YAML: %v", err)
	}
	
	log.Println("Export completed successfully.")
}
