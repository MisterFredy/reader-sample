package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func sendMail(to []string, cc []string, subject, message string) error {
	var err error

	err = godotenv.Load()

	body := "From: " + os.Getenv("EMAIL_SENDER") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", os.Getenv("EMAIL_SENDER"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST"))
	smtpAddr := fmt.Sprintf("%s:%d", os.Getenv("EMAIL_HOST"), os.Getenv("EMAIL_PORT"))

	err = smtp.SendMail(smtpAddr, auth, os.Getenv("EMAIL_SENDER"), append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
