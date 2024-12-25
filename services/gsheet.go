package services

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ibtikar-org-tr/iforgot/initializers"
	"google.golang.org/api/sheets/v4"
)

func GetGSheet(sheetID string) (*sheets.Spreadsheet, error) {
	if initializers.Srv == nil {
		return nil, errors.New("sheets service is not initialized")
	}

	// Open the spreadsheet using its ID and request grid data
	spreadsheet, err := initializers.Srv.Spreadsheets.Get(sheetID).IncludeGridData(true).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %w", err)
	}

	return spreadsheet, nil
}

func GetSheetTitle(sheetID string) (string, error) {
	sheet, err := GetGSheet(sheetID)
	if err != nil {
		return "", err
	}
	return sheet.Properties.Title, nil
}

// SearchValueInSheet searches for a specific value in a Google Sheet Page.
func SearchValueInSheet(valueToSearch string, valueColumn int, sheetID, pageName, lastColumn string) ([]interface{}, error) {
	if lastColumn == "" {
		lastColumn = "Z"
	}

	if initializers.Srv == nil {
		return nil, errors.New("sheets service is not initialized")
	}

	// readRange := "pageName!A:Z"
	readRange := fmt.Sprintf("%s!A:%s", pageName, lastColumn)
	resp, err := initializers.Srv.Spreadsheets.Values.Get(sheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}
	// fmt.Println(resp)

	// Iterate over the rows to search for the value in the second column
	found := false
	for i, row := range resp.Values {
		if len(row) > 0 && strings.Contains(fmt.Sprintf("%v", row[valueColumn]), valueToSearch) {
			// Print the entire row if the value matches in the second column
			fmt.Printf("Found in row %d: %v\n", i+2, row) // i+2 to adjust for the sheet row number
			found = true
			return row, nil
		}
	}

	// If no value was found, print a message
	if !found {
		fmt.Println("Value not found in the column.", valueColumn+1)
		return nil, nil
	}

	return nil, nil
}
