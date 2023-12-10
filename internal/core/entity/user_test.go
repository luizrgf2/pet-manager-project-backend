package entity_test

import (
	"testing"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	moks "github.com/luizrgf2/pet-manager-project-backend/internal/core/entity/moks"
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userMocked := moks.UserMock
	user, err := entity.NewUser(
		userMocked.NamePet,
		userMocked.Email,
		userMocked.Password,
		userMocked.AddrCep,
		userMocked.AddrStreet,
		userMocked.AddrCity,
		userMocked.AddrComplement,
		userMocked.AddrDistrict,
		userMocked.AddrNumber,
		userMocked.AddrState,
	)

	assert.Nil(t, err)
	assert.IsType(t, &entity.UserEntity{}, user)

}

func TestReturnErrorIfIsEmailInvalid(t *testing.T) {

	expectedError := errors.ErroBase{
		Message: errors.UserEmailInvalidErrorMessage,
		Code:    errors.UserEmailInvalidErrorCode,
	}

	userMocked := moks.UserMock
	_, err := entity.NewUser(
		userMocked.NamePet,
		"invalidEmail@",
		userMocked.Password,
		userMocked.AddrCep,
		userMocked.AddrStreet,
		userMocked.AddrCity,
		userMocked.AddrComplement,
		userMocked.AddrDistrict,
		userMocked.AddrNumber,
		userMocked.AddrState,
	)

	assert.Equal(t, expectedError.Error(), err.Error())

}

func TestReturnErrorIfIsEmaiIsEmpty(t *testing.T) {

	expectedError := errors.ErroBase{
		Message: errors.UserEmailInvalidErrorMessage,
		Code:    errors.UserEmailInvalidErrorCode,
	}

	userMocked := moks.UserMock
	_, err := entity.NewUser(
		userMocked.NamePet,
		"",
		userMocked.Password,
		userMocked.AddrCep,
		userMocked.AddrStreet,
		userMocked.AddrCity,
		userMocked.AddrComplement,
		userMocked.AddrDistrict,
		userMocked.AddrNumber,
		userMocked.AddrState,
	)

	assert.Equal(t, expectedError.Error(), err.Error())

}

func TestReturnErrorIfIsPassLenInvalid(t *testing.T) {

	expectedError := errors.ErroBase{
		Message: errors.UserPasswordLenErrorMessage,
		Code:    errors.UserPasswordUpperLetterErrorCode,
	}

	userMocked := moks.UserMock
	_, err := entity.NewUser(
		userMocked.NamePet,
		userMocked.Email,
		"123",
		userMocked.AddrCep,
		userMocked.AddrStreet,
		userMocked.AddrCity,
		userMocked.AddrComplement,
		userMocked.AddrDistrict,
		userMocked.AddrNumber,
		userMocked.AddrState,
	)

	assert.Equal(t, expectedError.Error(), err.Error())

}

func TestReturnErrorIfIsPassUpperLetterInvalid(t *testing.T) {

	expectedError := errors.ErroBase{
		Message: errors.UserPasswordUpperLetterErrorMessage,
		Code:    errors.UserPasswordUpperLetterErrorCode,
	}

	userMocked := moks.UserMock
	_, err := entity.NewUser(
		userMocked.NamePet,
		userMocked.Email,
		"luizfelipe",
		userMocked.AddrCep,
		userMocked.AddrStreet,
		userMocked.AddrCity,
		userMocked.AddrComplement,
		userMocked.AddrDistrict,
		userMocked.AddrNumber,
		userMocked.AddrState,
	)

	assert.Equal(t, expectedError.Error(), err.Error())

}

func TestReturnErrorIfCreateNewUserWithInvalidStateLocation(t *testing.T) {

	expectedError := errors.ErroBase{
		Message: errors.UserStateInvalidErrorMessage,
		Code:    errors.UserStateInvalidErrorCode,
	}

	userMocked := moks.UserMock
	_, err := entity.NewUser(
		userMocked.NamePet,
		userMocked.Email,
		userMocked.Password,
		userMocked.AddrCep,
		userMocked.AddrStreet,
		userMocked.AddrCity,
		userMocked.AddrComplement,
		userMocked.AddrDistrict,
		userMocked.AddrNumber,
		"mg",
	)

	assert.Equal(t, expectedError.Error(), err.Error())
}
