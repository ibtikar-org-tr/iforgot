package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ibtikar-org-tr/iforgot/initializers"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var srv *sheets.Service

func init() {
	// Load Env
	initializers.LoadEnv()

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
	srv, err = sheets.NewService(context.Background(), option.WithCredentialsJSON(data), option.WithScopes(scopes...))
	if err != nil {
		log.Fatalf("Unable to create Sheets client: %v\n", err)
	}

	fmt.Println("Sheets service initialized successfully")
}

func GetGSheet(sheetID string) *sheets.Spreadsheet {
	if srv == nil {
		fmt.Println("Sheets service is not initialized")
		return nil
	}

	// Open the spreadsheet using its ID
	spreadsheet, err := srv.Spreadsheets.Get(sheetID).Do()
	if err != nil {
		fmt.Printf("Unable to retrieve data from sheet: %v\n", err)
		return nil
	}

	return spreadsheet
}

func GetSheetTitle(sheetID string) string {
	sheet := GetGSheet(sheetID)
	if sheet == nil {
		return ""
	}
	return sheet.Properties.Title
}

func GetContactsPage(sheetID string) [][]interface{} {
	sheet := GetGSheet(sheetID)
	if sheet == nil {
		return nil
	}

	for _, sheet := range sheet.Sheets {
		if sheet.Properties.Title == "Contacts" {
			resp, err := srv.Spreadsheets.Values.Get(sheetID, sheet.Properties.Title).Do()
			if err != nil {
				fmt.Printf("Unable to retrieve data from sheet: %v\n", err)
				return nil
			}
			return resp.Values
		}
	}
	return nil
}

func GetSpecificContact(sheetID string, contactName string) map[string]interface{} {
	contactsPage := GetContactsPage(sheetID)
	if contactsPage == nil {
		return nil
	}

	for _, contact := range contactsPage {
		if contact[0] == contactName {
			contactDetails := make(map[string]interface{})
			for i, value := range contact {
				contactDetails[fmt.Sprintf("field%d", i)] = value
			}
			return contactDetails
		}
	}
	return nil
}
