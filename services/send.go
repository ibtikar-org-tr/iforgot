package services

import (
	"fmt"

	"github.com/ibtikar-org-tr/iforgot/initializers"
)

func SendResult(result []interface{}) error {
	mail_row := initializers.MailRow
	phone_row := initializers.PhoneRow
	number_row := initializers.NumberRow
	PasswordRow := initializers.PasswordRow

	// Check if indices are within bounds
	if mail_row >= len(result) || phone_row >= len(result) || number_row >= len(result) {
		return fmt.Errorf("index out of range")
	}

	mail_value := result[mail_row].(string)
	fmt.Print(mail_value)
	phone_value := result[phone_row].(string)
	fmt.Print(phone_value)
	number_value := result[number_row].(string)
	fmt.Print(number_value)
	password_value := result[PasswordRow].(string)
	fmt.Print(password_value)

	message := MessageStructure(mail_value, phone_value, number_value, password_value)

	go SendSMS(phone_value, message)
	emailRequest := EmailRequest{
		To:      mail_value,
		Subject: "iforgot - ibtikar assembly",
		Body:    message,
	}
	go SendMail(emailRequest)

	return nil
}

func MessageStructure(mail, phone, number, password string) string {
	header := initializers.HeaderMessage
	footer := initializers.FooterMessage
	message := header + "\n\nالبريد الإلكتروني: " + mail + "\nرقم الهاتف: " + phone + "\nرقم العضويّة: " + number + "\nكلمة المرور: " + password + "\n\n" + footer
	return message
}
