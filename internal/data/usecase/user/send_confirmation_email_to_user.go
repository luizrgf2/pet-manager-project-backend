package usecases

import (
	"strconv"

	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type SendConfirmationEmailToUserUseCase struct {
	jwtService                      services.JWTServiceInterface
	SMTPService                     services.SMTPService
	expirationTimeForTokenInSeconds uint
}

func (s *SendConfirmationEmailToUserUseCase) createTokenToSendWithEmail(idUser uint) (*string, error) {
	token, err := s.jwtService.CreateToken(strconv.Itoa(int(idUser)), s.expirationTimeForTokenInSeconds)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *SendConfirmationEmailToUserUseCase) sendEmailToUser(token string) error {
	err := s.SMTPService.SendConfirmationEmailToUser(token)
	if err != nil {
		return err
	}
	return nil
}

func (s SendConfirmationEmailToUserUseCase) Exec(input usecases.InputSendConfirmationEmailToSendUserUseCase) error {
	token, err := s.createTokenToSendWithEmail(input.IdUserToCreateToken)
	if err != nil {
		return err
	}

	err_to_send_email := s.sendEmailToUser(*token)
	if err_to_send_email != nil {
		return err_to_send_email
	}

	return nil
}
