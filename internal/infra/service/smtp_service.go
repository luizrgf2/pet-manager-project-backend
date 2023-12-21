package service

import (
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	"github.com/luizrgf2/pet-manager-project-backend/internal/infra/smtp"
)

type SMTPService struct {
}

func (s SMTPService) SendConfirmationEmailToUser(tokenOfConfirmation string, emailTo string) error {
	conn := smtp.NewSMTPConn()
	input := smtp.SMTPConfiguration{
		Message: "Confirme seu email " + tokenOfConfirmation,
		To:      emailTo,
		Subject: "Confirmação de email",
		Auth:    conn,
	}
	err := smtp.SendEmail(input)

	if err != nil {
		return &errors.ErroBase{
			Message: "Erro para enviar o email",
			Code:    500,
		}
	}
	return nil

}
