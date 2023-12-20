package repository

import (
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
)

type CreateUserRepositoryInput struct {
	NamePet        string
	Email          string
	Password       string
	AddrCep        string
	AddrStreet     string
	AddrNumber     uint
	AddrComplement *string
	AddrDistrict   string
	AddrCity       string
	AddrState      string
}

type UpdateUserRepositoryInput struct {
	Email          *string
	Password       *string
	NomePet        *string
	AddrCep        *string
	AddrStreet     *string
	AddrNumber     *uint
	AddrComplement *string
	AddrDistrict   *string
	AddrCity       *string
	AddrState      *string
}

type UserRepositoryInterface interface {
	Create(input CreateUserRepositoryInput) (*entity.UserEntity, error)
	FindById(id uint) (*entity.UserEntity, error)
	FindByEmail(email string) (*entity.UserEntity, error)
	Update(id uint, input UpdateUserRepositoryInput) (*entity.UserEntity, error)
	UpdateConfirmationToken(id uint, token string, expirationTimeInSeconds *uint) error
	Delete(id uint) error
}
