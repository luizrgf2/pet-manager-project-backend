package usecases

import (
	"strconv"

	core_error "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_error "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	"github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type SendConfirmationEmailToUserUseCase struct {
	JwtService                      services.JWTServiceInterface
	SMTPService                     services.SMTPServiceInterface
	ExpirationTimeForTokenInSeconds uint
	UserRepo                        repository.UserRepositoryInterface
}

func (s *SendConfirmationEmailToUserUseCase) createTokenToSendWithEmail(idUser uint) (*string, error) {
	token, err := s.JwtService.CreateToken(strconv.Itoa(int(idUser)), &s.ExpirationTimeForTokenInSeconds)
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

	confirmed, err := s.UserRepo.CheckIfUserConfirmed(input.IdUserToCreateToken)
	if err != nil {
		return err
	}

	if confirmed {
		return &core_error.ErroBase{
			Message: data_error.UserAlreadyConfirmedErrorMessage,
			Code:    data_error.UserAlreadyConfirmedErrorCode,
		}
	}

	token, err := s.createTokenToSendWithEmail(input.IdUserToCreateToken)
	if err != nil {
		return err
	}

	ErrToSendEmail := s.sendEmailToUser(*token)
	if ErrToSendEmail != nil {
		return ErrToSendEmail
	}

	return nil
}
