package entity

import (
	"time"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
)

var UserMock = entity.UserEntity{
	Id:             0,
	NamePet:        "ValidName",
	Email:          "validemail@gmail.com",
	Password:       "ValidPass123",
	AddrCep:        "38705280",
	AddrStreet:     "Alemar Rodrigues da cunha",
	AddrNumber:     622,
	AddrComplement: "",
	AddrDistrict:   "Sebastião Amorim",
	AddrCity:       "Patos de Minas",
	UpdatedAt:      time.Now(),
	CreatedAt:      time.Now(),
}
