package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ibtikar-org-tr/iforgot/initializers"
)

// Define a custom type for the allowed values
type ValueType string

const (
	Number ValueType = "number"
	Mail   ValueType = "mail"
	Phone  ValueType = "phone"
)

// Validate the typeOfValue
func validateTypeOfValue(typeOfValue string) (ValueType, error) {
	switch ValueType(typeOfValue) {
	case Number, Mail, Phone:
		return ValueType(typeOfValue), nil
	default:
		return "", errors.New("invalid typeOfValue, must be one of: number, mail, phone")
	}
}

func SearchMain(value, typeOfValue, ip string) (string, error) {
	// Check if the session is valid
	session, _ := GetSessionByIP(ip)
	if session != nil {
		session_police := CheckLastSent(session)
		if !session_police {
			return "", errors.New("too many requests")
		}
	}

	// Declarations
	value_min := initializers.ValueMin
	value_max := initializers.ValueMax

	sheetID := initializers.SheetID
	pageName := initializers.PageName
	lastColumn := initializers.LastColumn

	// Validate the value
	validTypeOfValue, err := validateTypeOfValue(typeOfValue)
	if err != nil {
		return "", err
	}
	var valueColumn int
	switch validTypeOfValue {
	case Mail:
		valueColumn = initializers.MailRow
		if !IsValidEmail(value) {
			return "", errors.New("invalid email format")
		}
	case Phone:
		valueColumn = initializers.PhoneRow
		if !IsValidPhone(value) {
			return "", errors.New("invalid phone format")
		}
	case Number:
		valueColumn = initializers.NumberRow
		number, err := strconv.Atoi(value)
		if err != nil {
			return "", errors.New("invalid number format")
		}
		if number < value_min || number > value_max {
			return "", errors.New("number out of range")
		}
	default:
		return "", errors.New("invalid type")
	}

	// Log the parameters
	fmt.Printf("Searching for value: %s, type: %s, sheetID: %s, pageName: %s, lastColumn: %s\n", value, typeOfValue, sheetID, pageName, lastColumn)

	// Search the value
	result, err := SearchValueInSheet(value, valueColumn, sheetID, pageName, lastColumn)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve data from sheet: %w", err)
	}

	// Send the result
	go SendResult(result)

	// Create a new session
	go StoreSession(ip)

	return "", nil
}
