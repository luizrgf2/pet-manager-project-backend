package user_test

import (
	"testing"

	moks "github.com/luizrgf2/pet-manager-project-backend/internal/core/entity/moks"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	repository "github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
	services_mocked "github.com/luizrgf2/pet-manager-project-backend/test/moks/service"
	"github.com/stretchr/testify/assert"
)

var cepServiceMocked = services_mocked.CEPServiceInMemory{}
var hashServiceMocked = services_mocked.HashServiceInMemory{}
var userRepo = repository.UserRepository{}

var sut = usecases.CreateUserUseCase{
	CepService:     cepServiceMocked,
	HashService:    hashServiceMocked,
	UserRepository: userRepo,
}

func TestCreateUse(t *testing.T) {

	userToTest := moks.UserMock
	input := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        userToTest.NamePet,
		Email:          userToTest.Email,
		Password:       userToTest.Password,
		AddrCep:        userToTest.AddrCep,
		AddrComplement: userToTest.AddrComplement,
		AddrNumber:     userToTest.AddrNumber,
	}
	_, err := sut.Exec(input)

	assert.Nil(t, err)

}
