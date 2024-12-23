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

func GetContactsPage(sheetID string) [][]interface{} {
	sheet := GetGSheet(sheetID)
	if sheet == nil {
		return nil
	}

	for _, sheet := range sheet.Sheets {
		if sheet.Properties.Title == "Contacts" {
			resp, err := initializers.Srv.Spreadsheets.Values.Get(sheetID, sheet.Properties.Title).Do()
			if err != nil {
				fmt.Printf("Unable to retrieve data from sheet: %v\n", err)
				return nil
			}
			return resp.Values
		}
	}
	return nil
}
