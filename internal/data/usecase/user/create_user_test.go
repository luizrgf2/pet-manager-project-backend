package usecases_test

import (
	"testing"

	errors_core "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	errors_data "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	usecases_imp "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	repository_moks "github.com/luizrgf2/pet-manager-project-backend/test/moks/repository"
	tests_moks "github.com/luizrgf2/pet-manager-project-backend/test/moks/service"

	"github.com/stretchr/testify/assert"
)

var cepService = tests_moks.CEPServiceInMemory{}
var hashService = tests_moks.HashServiceInMemory{}
var userRepo = repository_moks.UserRepositoryInMemory{}
var jwtService = tests_moks.JWTServiceInMemory{}

var sut = usecases_imp.CreateUserUseCase{
	CepService:     cepService,
	UserRepository: userRepo,
	HashService:    hashService,
}

func TestCreateUserUseCase(t *testing.T) {

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid.com",
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	_, err := sut.Exec(userToTeste)

	assert.Nil(t, err)
}

func TestReturnErrorWIthInvalidEmail(t *testing.T) {

	expectedError := errors_core.ErroBase{
		Message: errors_core.UserEmailInvalidErrorMessage,
		Code:    errors_core.UserNameInvalidErrorCode,
	}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid",
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	_, err := sut.Exec(userToTeste)

	assert.Equal(t, expectedError.Error(), err.Error())
}

func TestReturnErrorWIthInvalidCep(t *testing.T) {

	expectedError := errors_core.ErroBase{
		Message: errors_data.CEPInvalidErrorMessage,
		Code:    uint(errors_data.CEPInvalidErrorCode),
	}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid.com",
		Password:       "Teste12345",
		AddrCep:        "234",
		AddrComplement: "",
		AddrNumber:     622,
	}

	_, err := sut.Exec(userToTeste)

	assert.Equal(t, expectedError.Error(), err.Error())
}

func TestReturnErrorWIthInvalidPasswordLen(t *testing.T) {

	expectedError := errors_core.ErroBase{
		Message: errors_core.UserPasswordLenErrorMessage,
		Code:    errors_core.UserPasswordLenErrorCode,
	}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid.com",
		Password:       "1234",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	_, err := sut.Exec(userToTeste)

	assert.Equal(t, expectedError.Error(), err.Error())
}

func TestReturnErrorWIthInvalidPasswordUpperCase(t *testing.T) {

	expectedError := errors_core.ErroBase{
		Message: errors_core.UserPasswordUpperLetterErrorMessage,
		Code:    errors_core.UserPasswordUpperLetterErrorCode,
	}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid.com",
		Password:       "teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	_, err := sut.Exec(userToTeste)

	assert.Equal(t, expectedError.Error(), err.Error())
}

func TestReturnErrorIfTryCreateUserWithEmailAlreadyExists(t *testing.T) {

	expectedError := errors_core.ErroBase{
		Message: errors_core.UserAlreadyExistsErrorMessage,
		Code:    errors_core.UserAlreadyExistsErrorCode,
	}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid1.com",
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	_, err := sut.Exec(userToTeste)

	assert.Equal(t, expectedError.Error(), err.Error())
}
