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
	SMTP_User  string
	SMTP_Pass  string
	SMTP_Host  string
	SMTP_Port  string
	SMS_MS     string
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

	SMTP_User := os.Getenv("SMTP_USER")
	if SMTP_User == "" {
		log.Fatal("SMTP_USER environment variable is not set")
	}

	SMTP_Pass := os.Getenv("SMTP_PASS")
	if SMTP_Pass == "" {
		log.Fatal("SMTP_PASS environment variable is not set")
	}

	SMTP_Host := os.Getenv("SMTP_HOST")
	if SMTP_Host == "" {
		log.Fatal("SMTP_HOST environment variable is not set")
	}

	SMTP_Port := os.Getenv("SMTP_PORT")
	if SMTP_Port == "" {
		log.Fatal("SMTP_PORT environment variable is not set")
	}

	SMS_MS := os.Getenv("SMS_MS")
	if SMS_MS == "" {
		log.Fatal("SMS_MS environment variable is not set")
	}
}
