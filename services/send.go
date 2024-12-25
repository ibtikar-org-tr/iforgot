package services

import (
	"fmt"

	"github.com/ibtikar-org-tr/iforgot/initializers"
)

func SendResult(result []interface{}) error {
	mail_row := initializers.MailRow
	phone_row := initializers.PhoneRow
	number_row := initializers.NumberRow

	mail_value := result[mail_row].(string)
	fmt.Print(mail_value)
	phone_value := result[phone_row].(string)
	fmt.Print(phone_value)
	number_value := result[number_row].(string)
	fmt.Print(number_value)

	message := MessageStructure(mail_value, phone_value, number_value)

	go SendSMS(phone_value, message)
	emailRequest := EmailRequest{
		To:      mail_value,
		Subject: "Result Notification",
		Body:    message,
	}
	go SendMail(emailRequest)

	return nil
}

func MessageStructure(mail, phone, number string) string {
	message := "البريد الإلكتروني: " + mail + "\nرقم الهاتف: " + phone + "\nرقم العضويّة: " + number
	return message
}
