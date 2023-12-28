package user_usercases

import (
	"strconv"

	core_error "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	user_usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
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

func (s *SendConfirmationEmailToUserUseCase) sendEmailToUser(token string, emailTo string) error {
	err := s.SMTPService.SendConfirmationEmailToUser(token, emailTo)
	if err != nil {
		return err
	}
	return nil
}

func (s SendConfirmationEmailToUserUseCase) Exec(input user_usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase) error {

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

	user, err := s.UserRepo.FindById(input.IdUserToCreateToken)
	if err != nil {
		return err
	}

	token, err := s.createTokenToSendWithEmail(input.IdUserToCreateToken)
	if err != nil {
		return err
	}

	errUpdasteTokenInsideUser := s.UserRepo.UpdateConfirmationToken(user.Id, *token, &s.ExpirationTimeForTokenInSeconds)
	if errUpdasteTokenInsideUser != nil {
		return errUpdasteTokenInsideUser
	}

	ErrToSendEmail := s.sendEmailToUser(*token, user.Email)
	if ErrToSendEmail != nil {
		return ErrToSendEmail
	}

	return nil
}
