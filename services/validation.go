package services

import (
	"errors"
	"regexp"
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

// isValidPhone validates the phone number format
func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(phone)
}

// isValidEmail validates the email format
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
