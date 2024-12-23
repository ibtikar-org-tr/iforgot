package services

import (
	"fmt"
	"net/http"
	"os"
)

func SendSMS(phone string, message string) {
	smsMs := os.Getenv("SMS_MS")
	url := fmt.Sprintf("%s?phone=%s&message=%s", smsMs, phone, message)

	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Printf("Failed to send SMS: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to send SMS: %s\n", resp.Status)
		return
	}

	fmt.Printf("SMS sent successfully to %s\n", phone)
}
