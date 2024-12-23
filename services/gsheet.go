package services

import (
	"fmt"

	"github.com/ibtikar-org-tr/iforgot/initializers"
	"google.golang.org/api/sheets/v4"
)

func GetGSheet(sheetID string) *sheets.Spreadsheet {
	if initializers.Srv == nil {
		fmt.Println("Sheets service is not initialized")
		return nil
	}

	// Open the spreadsheet using its ID
	spreadsheet, err := initializers.Srv.Spreadsheets.Get(sheetID).Do()
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
