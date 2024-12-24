package services

import (
	"errors"
	"os"
	"strconv"
)

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
	value_min, err := strconv.Atoi(os.Getenv("VALUE_MIN"))
	if err != nil {
		return "", errors.New("invalid VALUE_MIN environment variable")
	}
	value_max, err := strconv.Atoi(os.Getenv("VALUE_MAX"))
	if err != nil {
		return "", errors.New("invalid VALUE_MAX environment variable")
	}
	sheetID := os.Getenv("SHEET_ID")
	pageName := os.Getenv("PAGE_NAME")
	lastColumn := os.Getenv("LAST_COLUMN")

	// Validate the value
	if typeOfValue == "email" {
		if !IsValidEmail(value) {
			return "", errors.New("invalid email format")
		}
	} else if typeOfValue == "phone" {
		if !IsValidPhone(value) {
			return "", errors.New("invalid phone format")
		}
	} else if typeOfValue == "number" {
		number, err := strconv.Atoi(value)
		if err != nil {
			return "", errors.New("invalid number format")
		}
		if number < value_min || number > value_max {
			return "", errors.New("number out of range")
		}
	} else {
		return "", errors.New("invalid type")
	}

	// Search the value
	result, err := SearchValueInSheet(value, sheetID, pageName, lastColumn)
	if err != nil {
		return "", err
	}

	// Send the result
	SendResult(result)

	// Create a new session
	err = StoreSession(ip)
	if err != nil {
		return "", err
	}

	return "", nil
}
