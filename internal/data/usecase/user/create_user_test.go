package usecases_test

import (
	"testing"

	errors_core "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	errors_data "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	usecases_imp "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	tests_moks "github.com/luizrgf2/pet-manager-project-backend/test/moks/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase(t *testing.T) {

	cepService := tests_moks.CEPServiceInMemory{}

	sut := usecases_imp.CreateUserUseCase{CepService: cepService}

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

	cepService := tests_moks.CEPServiceInMemory{}
	sut := usecases_imp.CreateUserUseCase{CepService: cepService}

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

	cepService := tests_moks.CEPServiceInMemory{}
	sut := usecases_imp.CreateUserUseCase{CepService: cepService}

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

	cepService := tests_moks.CEPServiceInMemory{}
	sut := usecases_imp.CreateUserUseCase{CepService: cepService}

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

	cepService := tests_moks.CEPServiceInMemory{}
	sut := usecases_imp.CreateUserUseCase{CepService: cepService}

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
