package services

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/ibtikar-org-tr/iforgot/initializers"
)

type EmailRequest struct {
	To      string `json:"to" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func SendMail(emailReq EmailRequest) error {
	from := initializers.SMTP_User
	password := initializers.SMTP_Pass
	smtpHost := initializers.SMTP_Host
	smtpPort := initializers.SMTP_Port

	if from == "" || password == "" || smtpHost == "" || smtpPort == "" {
		log.Println("SMTP environment variables not set")
		return fmt.Errorf("SMTP environment variables not set")
	}

	recipients := strings.Split(emailReq.To, ",")
	// Filter out empty email addresses
	var validRecipients []string
	for _, recipient := range recipients {
		if strings.TrimSpace(recipient) != "" {
			validRecipients = append(validRecipients, recipient)
		}
	}

	// Log the recipient email addresses
	log.Printf("Sending email to: %v", validRecipients)

	if len(validRecipients) == 0 {
		log.Println("No valid email addresses to send to.")
		return nil
	}

	message := []byte("Subject: " + emailReq.Subject + "\r\n\r\n" + emailReq.Body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, validRecipients, message)
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		return err
	}

	log.Println("Mail sent successfully")
	return nil
}
