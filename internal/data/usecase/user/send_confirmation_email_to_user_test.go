package user_usercases_test

import (
	"testing"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	user_usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"

	data_errors "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	"github.com/stretchr/testify/assert"

	repository "github.com/luizrgf2/pet-manager-project-backend/test/moks/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/test/moks/service"
)

func TestSendConfirmationEmail(t *testing.T) {
	jwtService := services.JWTServiceInMemory{}
	smtpService := services.SMTPServiceInMemory{}
	userRepo := repository.UserRepositoryInMemory{}

	sut := usecases.SendConfirmationEmailToUserUseCase{
		JwtService:                      jwtService,
		SMTPService:                     smtpService,
		ExpirationTimeForTokenInSeconds: 21600,
		UserRepo:                        userRepo,
	}
	input := user_usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 1,
	}
	err := sut.Exec(input)
	assert.Nil(t, err)
}

func TestReturnErrorIfSendEmailToUserConfirmed(t *testing.T) {

	expectedError := &core_errors.ErroBase{
		Message: data_errors.UserAlreadyConfirmedErrorMessage,
		Code:    data_errors.UserAlreadyConfirmedErrorCode,
	}

	jwtService := services.JWTServiceInMemory{}
	smtpService := services.SMTPServiceInMemory{}
	userRepo := repository.UserRepositoryInMemory{}

	sut := usecases.SendConfirmationEmailToUserUseCase{
		JwtService:                      jwtService,
		SMTPService:                     smtpService,
		ExpirationTimeForTokenInSeconds: 21600,
		UserRepo:                        userRepo,
	}
	input := user_usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 2,
	}
	err := sut.Exec(input)
	assert.Equal(t, expectedError, err)
}

func TestReturnErrorIfSendEmailToUserNotExists(t *testing.T) {

	expectedError := &core_errors.ErroBase{
		Message: core_errors.UserNotExistsErrorMessage,
		Code:    core_errors.UserNotExistsErrorCode,
	}

	jwtService := services.JWTServiceInMemory{}
	smtpService := services.SMTPServiceInMemory{}
	userRepo := repository.UserRepositoryInMemory{}

	sut := usecases.SendConfirmationEmailToUserUseCase{
		JwtService:                      jwtService,
		SMTPService:                     smtpService,
		ExpirationTimeForTokenInSeconds: 21600,
		UserRepo:                        userRepo,
	}
	input := user_usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 3,
	}
	err := sut.Exec(input)
	assert.Equal(t, expectedError, err)
}
