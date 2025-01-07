package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ibtikar-org-tr/iforgot/initializers"
)

func SearchMain(value, typeOfValue, ip string) (string, error) {
	// Check if the session is valid
	session, _ := GetSessionByIP(ip)
	if session != nil {
		session_police := CheckLastSent(session)
		if !session_police {
			return "err: too many requests", errors.New("too many requests")
		}
	}

	// Log the request
	fmt.Printf("Searching for value: %s, type: %s\n", value, typeOfValue)

	// Declarations
	value_min := initializers.ValueMin
	value_max := initializers.ValueMax

	sheetID := initializers.SheetID
	pageName := initializers.PageName
	lastColumn := initializers.LastColumn

	// Validate the value
	validTypeOfValue, err := validateTypeOfValue(typeOfValue)
	if err != nil {
		return "err: validation error", err
	}
	var valueColumn int
	switch validTypeOfValue {
	case Mail:
		valueColumn = initializers.MailRow
		if !IsValidEmail(value) {
			return "err: invalid email format", errors.New("invalid email format")
		}
	case Phone:
		valueColumn = initializers.PhoneRow
		// if !IsValidPhone(value) {
		// 	return "", errors.New("invalid phone format")
		// }
	case Number:
		valueColumn = initializers.NumberRow
		number, err := strconv.Atoi(value)
		if err != nil {
			return "err: invalid number format", errors.New("invalid number format")
		}
		if number < value_min || number > value_max {
			return "err: number out of range", errors.New("number out of range")
		}
	default:
		return "err: invalid type", errors.New("invalid type")
	}

	// Log the parameters
	fmt.Printf("Searching for value: %s, type: %s, sheetID: %s, pageName: %s, lastColumn: %s\n", value, typeOfValue, sheetID, pageName, lastColumn)

	// Search the value
	result, err := SearchValueInSheet(value, valueColumn, sheetID, pageName, lastColumn)
	if err != nil {
		return "err: unable to retrieve data", fmt.Errorf("unable to retrieve data from sheet: %w", err)
	}

	// Check if result is empty
	if result == nil {
		return "err: value not found", errors.New("value not found")
	}

	// Send the result
	go SendResult(result)

	// Create a new session
	go StoreSession(ip)

	mail := result[initializers.MailRow].(string)

	return mail, nil
}
