package services

import (
	"fmt"
	"os"
	"strconv"
)

func SendResult(result []interface{}) error {
	mail_row, err := strconv.Atoi(os.Getenv("MAIL_ROW"))
	if err != nil {
		return err
	}
	phone_row, err := strconv.Atoi(os.Getenv("PHONE_ROW"))
	if err != nil {
		return err
	}
	number_row, err := strconv.Atoi(os.Getenv("NUMBER_ROW"))
	if err != nil {
		return err
	}

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
