package initializers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MailRow       int
	PhoneRow      int
	NumberRow     int
	PasswordRow   int
	ValueMin      int
	ValueMax      int
	HeaderMessage string
	FooterMessage string
	SheetID       string
	PageName      string
	LastColumn    string
	SMTP_User     string
	SMTP_Pass     string
	SMTP_Host     string
	SMTP_Port     string
	SMS_MS        string
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
		fmt.Printf("Invalid MAIL_ROW environment variable: %v", err)
	}

	PhoneRow, err = strconv.Atoi(os.Getenv("PHONE_ROW"))
	if err != nil {
		fmt.Printf("Invalid PHONE_ROW environment variable: %v", err)
	}

	NumberRow, err = strconv.Atoi(os.Getenv("NUMBER_ROW"))
	if err != nil {
		fmt.Printf("Invalid NUMBER_ROW environment variable: %v", err)
	}

	PasswordRow, err = strconv.Atoi(os.Getenv("PASSWORD_ROW"))
	if err != nil {
		fmt.Printf("Invalid PASSWORD_ROW environment variable: %v", err)
	}

	ValueMin, err = strconv.Atoi(os.Getenv("VALUE_MIN"))
	if err != nil {
		fmt.Printf("Invalid VALUE_MIN environment variable: %v", err)
	}

	ValueMax, err = strconv.Atoi(os.Getenv("VALUE_MAX"))
	if err != nil {
		fmt.Printf("Invalid VALUE_MAX environment variable: %v", err)
	}

	HeaderMessage = os.Getenv("HEADER_MESSAGE")
	if HeaderMessage == "" {
		fmt.Printf("HEADER_MESSAGE environment variable is not set")
	}

	FooterMessage = os.Getenv("FOOTER_MESSAGE")
	if FooterMessage == "" {
		fmt.Printf("FOOTER_MESSAGE environment variable is not set")
	}

	SheetID = os.Getenv("SHEET_ID")
	if SheetID == "" {
		fmt.Printf("SHEET_ID environment variable is not set")
	}

	PageName = os.Getenv("PAGE_NAME")
	if PageName == "" {
		fmt.Printf("PAGE_NAME environment variable is not set")
	}

	LastColumn = os.Getenv("LAST_COLUMN")
	if LastColumn == "" {
		fmt.Printf("LAST_COLUMN environment variable is not set")
	}

	SMTP_User = os.Getenv("SMTP_USER")
	if SMTP_User == "" {
		fmt.Printf("SMTP_USER environment variable is not set")
	}

	SMTP_Pass = os.Getenv("SMTP_PASS")
	if SMTP_Pass == "" {
		fmt.Printf("SMTP_PASS environment variable is not set")
	}

	SMTP_Host = os.Getenv("SMTP_HOST")
	if SMTP_Host == "" {
		fmt.Printf("SMTP_HOST environment variable is not set")
	}

	SMTP_Port = os.Getenv("SMTP_PORT")
	if SMTP_Port == "" {
		fmt.Printf("SMTP_PORT environment variable is not set")
	}

	SMS_MS = os.Getenv("SMS_MS")
	if SMS_MS == "" {
		fmt.Printf("SMS_MS environment variable is not set")
	}
}
