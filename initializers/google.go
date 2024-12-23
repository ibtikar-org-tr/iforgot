package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var Srv *sheets.Service

func SrvInit() {
	// Path to your service account key file
	serviceAccountFile := os.Getenv("GOOGLE_API_JSON_PATH")
	if serviceAccountFile == "" {
		log.Fatalf("GOOGLE_API_JSON_PATH environment variable is not set")
	}

	// Normalize the file path
	serviceAccountFile = filepath.Clean(serviceAccountFile)

	// Read the service account JSON file
	data, err := os.ReadFile(serviceAccountFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v\n", err)
	}

	// Define required scopes
	scopes := []string{
		"https://www.googleapis.com/auth/spreadsheets",
		"https://www.googleapis.com/auth/drive",
	}

	// Initialize the Sheets service
	Srv, err = sheets.NewService(context.Background(), option.WithCredentialsJSON(data), option.WithScopes(scopes...))
	if err != nil {
		log.Fatalf("Unable to create Sheets client: %v\n", err)
	}

	fmt.Println("Sheets service initialized successfully")
}
