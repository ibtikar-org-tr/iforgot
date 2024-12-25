package initializers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MailRow    int
	PhoneRow   int
	NumberRow  int
	ValueMin   int
	ValueMax   int
	SheetID    string
	PageName   string
	LastColumn string
)

func LoadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse and assign environment variables to global variables
	MailRow, err = strconv.Atoi(os.Getenv("MAIL_ROW"))
	if err != nil {
		log.Fatalf("Invalid MAIL_ROW environment variable: %v", err)
	}

	PhoneRow, err = strconv.Atoi(os.Getenv("PHONE_ROW"))
	if err != nil {
		log.Fatalf("Invalid PHONE_ROW environment variable: %v", err)
	}

	NumberRow, err = strconv.Atoi(os.Getenv("NUMBER_ROW"))
	if err != nil {
		log.Fatalf("Invalid NUMBER_ROW environment variable: %v", err)
	}

	ValueMin, err = strconv.Atoi(os.Getenv("VALUE_MIN"))
	if err != nil {
		log.Fatalf("Invalid VALUE_MIN environment variable: %v", err)
	}

	ValueMax, err = strconv.Atoi(os.Getenv("VALUE_MAX"))
	if err != nil {
		log.Fatalf("Invalid VALUE_MAX environment variable: %v", err)
	}

	SheetID = os.Getenv("SHEET_ID")
	if SheetID == "" {
		log.Fatal("SHEET_ID environment variable is not set")
	}

	PageName = os.Getenv("PAGE_NAME")
	if PageName == "" {
		log.Fatal("PAGE_NAME environment variable is not set")
	}

	LastColumn = os.Getenv("LAST_COLUMN")
	if LastColumn == "" {
		log.Fatal("LAST_COLUMN environment variable is not set")
	}
}
