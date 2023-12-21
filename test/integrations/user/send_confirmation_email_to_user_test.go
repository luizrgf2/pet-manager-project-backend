package user_test

import (
	"fmt"
	"testing"

	"github.com/luizrgf2/pet-manager-project-backend/config"
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_errors "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	repository "github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	services_moks "github.com/luizrgf2/pet-manager-project-backend/test/moks/service"
	"github.com/stretchr/testify/assert"
)

func init() {
	cepService := services_moks.CEPServiceInMemory{}
	hashService := services.HashService{}
	userRepo := repository.UserRepository{}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          config.SMTP_EMAIL_RECEIVER_TO_TEST,
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	createUserUseCase := usecases.CreateUserUseCase{
		CepService:     cepService,
		UserRepository: userRepo,
		HashService:    hashService,
	}

	_, err := createUserUseCase.Exec(userToTeste)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("users created!!")
	}

}

func TestSendConfirmationEmail(t *testing.T) {
	jwtService := services.JWTService{}
	smtpService := services.SMTPService{}
	userRepo := repository.UserRepository{}

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

	jwtService := services.JWTService{}
	smtpService := services.SMTPService{}
	userRepo := repository.UserRepository{}

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

	jwtService := services.JWTService{}
	smtpService := services.SMTPService{}
	userRepo := repository.UserRepository{}

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
