package usecases_test

import (
	"testing"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
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
	input := usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 1,
	}
	err := sut.Exec(input)
	assert.Nil(t, err)
}

func TestReturnErrorIfSendEmailToUserConfirmed(t *testing.T) {

	expectedError := &errors.ErroBase{
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
	input := usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 2,
	}
	err := sut.Exec(input)
	assert.Equal(t, expectedError, err)
}

func TestReturnErrorIfSendEmailToUserNotExists(t *testing.T) {

	expectedError := &errors.ErroBase{
		Message: errors.UserNotExistsErrorMessage,
		Code:    errors.UserNotExistsErrorCode,
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
	input := usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 3,
	}
	err := sut.Exec(input)
	assert.Equal(t, expectedError, err)
}
