package test

import (
	"time"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	"github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
)

type UserRepositoryInMemory struct {
	users []entity.UserEntity
}

func (U UserRepositoryInMemory) Create(input repository.CreateUserRepositoryInput) (*entity.UserEntity, error) {

	var id uint = 1

	if len(U.users) > 0 {
		lenOfUsers := len(U.users)
		lastUser := U.users[lenOfUsers-1]
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

	U.users = append(U.users, user)

	return &user, nil

}

func (U UserRepositoryInMemory) FindById(id uint) (*entity.UserEntity, error) {
	for _, user := range U.users {
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
	for _, user := range U.users {
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

	for _, usr := range U.users {
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

	for _, usr := range U.users {
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

	for _, usr := range U.users {
		if usr.Id != id {
			U.users = append(U.users, usr)
		}
	}
	return nil
}
