package service

import (
	"strings"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	infra_error "github.com/luizrgf2/pet-manager-project-backend/internal/infra/error/services"
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
		if strings.Contains(err.Error(), "dial tcp: lookup") && strings.Contains(err.Error(), "no such host") {
			return &core_errors.ErroBase{
				Message: infra_error.EmailNotFoundedErrorMessage,
				Code:    infra_error.EmailNotFoundedErrorCode,
			}
		}

		return &core_errors.ErroBase{
			Message: "Erro para enviar o email",
			Code:    500,
		}
	}
	return nil

}
