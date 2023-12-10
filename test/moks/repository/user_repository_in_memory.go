package test

import (
	"time"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	"github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
)

type UserRepositoryInMemory struct {
	Users []entity.UserEntity
}

func (U UserRepositoryInMemory) Create(input repository.CreateUserRepositoryInput) (*entity.UserEntity, error) {

	var id uint = 1

	if len(U.Users) > 0 {
		lenOfUsers := len(U.Users)
		lastUser := U.Users[lenOfUsers-1]
		id = uint(lastUser.Id) + 1
	}

	user := entity.UserEntity{
		Id:             id,
		NamePet:        input.NamePet,
		Email:          input.Email,
		Password:       input.Password,
		AddrCep:        input.AddrCep,
		AddrStreet:     input.AddrStreet,
		AddrNumber:     input.AddrNumber,
		AddrComplement: *input.AddrComplement,
		AddrDistrict:   input.AddrDistrict,
		AddrCity:       input.AddrCity,
		AddrState:      input.AddrState,
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	U.Users = append(U.Users, user)

	return &user, nil

}

func (U UserRepositoryInMemory) FindById(id uint) (*entity.UserEntity, error) {

	userToTest := entity.UserEntity{
		Id:             20,
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid1.com",
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
		AddrStreet:     "Alemar Rodrigues da Cunha",
		AddrDistrict:   "Sebastião Amotim",
		AddrCity:       "Patos de Minas",
		AddrState:      "MG",
	}

	U.Users = append(U.Users, userToTest)

	for _, user := range U.Users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, &errors.ErroBase{
		Message: errors.UserNotExistsErrorMessage,
		Code:    errors.UserNotExistsErrorCode,
	}
}

func (U UserRepositoryInMemory) FindByEmail(email string) (*entity.UserEntity, error) {

	userToTest := entity.UserEntity{
		Id:             1,
		NamePet:        "Felicidog pet salon",
		Email:          "email@valid1.com",
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
		AddrStreet:     "Alemar Rodrigues da Cunha",
		AddrDistrict:   "Sebastião Amotim",
		AddrCity:       "Patos de Minas",
		AddrState:      "MG",
	}

	U.Users = append(U.Users, userToTest)

	for _, user := range U.Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, &errors.ErroBase{
		Message: errors.UserNotExistsErrorMessage,
		Code:    errors.UserNotExistsErrorCode,
	}
}

func (U UserRepositoryInMemory) Update(id uint, input repository.UpdateUserRepositoryInput) (*entity.UserEntity, error) {
	var user *entity.UserEntity = nil

	for _, usr := range U.Users {
		if usr.Id == id {
			user = &usr
		}
	}

	if user == nil {
		return nil, &errors.ErroBase{
			Message: errors.UserNotExistsErrorMessage,
			Code:    errors.UserNotExistsErrorCode,
		}
	}

	if input.AddrCep != nil {
		user.AddrCep = *input.AddrCep
	}

	if input.AddrCity != nil {
		user.AddrCity = *input.AddrCity
	}

	if input.AddrComplement != nil {
		user.AddrComplement = *input.AddrComplement
	}

	if input.AddrDistrict != nil {
		user.AddrDistrict = *input.AddrDistrict
	}

	if input.AddrNumber != nil {
		user.AddrNumber = *input.AddrNumber
	}

	if input.AddrState != nil {
		user.AddrState = *input.AddrState
	}

	if input.AddrStreet != nil {
		user.AddrStreet = *input.AddrStreet
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.NomePet != nil {
		user.NamePet = *input.NomePet
	}

	if input.Password != nil {
		user.Password = *input.Password
	}

	return user, nil

}

func (U UserRepositoryInMemory) Delete(id uint) error {
	var user *entity.UserEntity = nil

	for _, usr := range U.Users {
		if usr.Id == id {
			user = &usr
		}
	}

	if user == nil {
		return &errors.ErroBase{
			Message: errors.UserNotExistsErrorMessage,
			Code:    errors.UserNotExistsErrorCode,
		}
	}

	for _, usr := range U.Users {
		if usr.Id != id {
			U.Users = append(U.Users, usr)
		}
	}
	return nil
}
