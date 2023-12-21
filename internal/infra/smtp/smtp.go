package smtp

import (
	"fmt"
	"net/smtp"

	"github.com/luizrgf2/pet-manager-project-backend/config"
)

type SMTPConfiguration struct {
	Message string
	To      string
	Subject string
	Auth    smtp.Auth
}

type SMTPConn struct {
	port uint
}

func NewSMTPConn() smtp.Auth {
	username := config.SMTP_EMAIL_SENDER
	password := config.SMTP_PASSWORD
	host := config.SMTP_SERVER

	auth := smtp.PlainAuth("", username, password, host)
	return auth
}

func SendEmail(input SMTPConfiguration) error {
	To := input.To
	to := []string{To}
	port := uint(config.SMTP_PORT)

	message := fmt.Sprintf("To: %s\r\n"+

		"Subject: %s\r\n"+

		"\r\n"+

		"%s\r\n", To, input.Subject, input.Message)

	errToSendEmail := smtp.SendMail(fmt.Sprintf("%s:%d", config.SMTP_SERVER, port), input.Auth, config.SMTP_EMAIL_SENDER, to, []byte(message))

	if errToSendEmail != nil {
		return errToSendEmail
	}
	return nil
}
