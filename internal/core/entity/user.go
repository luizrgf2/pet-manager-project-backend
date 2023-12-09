package entity

import (
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
)

type UserEntity struct {
	Id             uint
	NamePet        string
	Email          string
	Password       string
	AddrCep        string
	AddrStreet     string
	AddrNumber     uint
	AddrComplement string
	AddrDistrict   string
	AddrCity       string
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

func (U *UserEntity) IsValidName() bool {
	nameLen := utf8.RuneCountInString(U.NamePet)
	if nameLen < 4 || nameLen > 50 {
		return false
	}
	return true
}

func (U *UserEntity) IsValidEmail() bool {
	regexEmail := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validEmail, err := regexp.Match(regexEmail, []byte(U.Email))
	if err != nil {
		return false
	}
	return validEmail
}

func (U *UserEntity) IsValidLenPassword() bool {
	passLen := utf8.RuneCountInString(U.Password)
	if passLen < 8 || passLen > 15 {
		return false
	}
	return true
}

func (U *UserEntity) IsValidUpperLetterPassword() bool {
	regexPassUpperLetter := `[A-Z]`
	validPass, err := regexp.Match(regexPassUpperLetter, []byte(U.Password))
	if err != nil {
		return false
	}
	return validPass
}

func NewUser(NamePet string, Email string, Password string, AddrCep string, AddrStreet string, AddrCity string, AddrComplement string, AddrDistrict string, AddrNumber uint) (*UserEntity, error) {
	user := &UserEntity{
		Id:             0,
		NamePet:        NamePet,
		Email:          Email,
		Password:       Password,
		AddrCep:        AddrCep,
		AddrStreet:     AddrStreet,
		AddrNumber:     AddrNumber,
		AddrComplement: AddrComplement,
		AddrDistrict:   AddrDistrict,
		AddrCity:       AddrCity,
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	if !user.IsValidEmail() {
		return nil, &errors.ErroBase{Message: errors.UserEmailInvalidErrorMessage, Code: errors.UserEmailInvalidErrorCode}
	}

	if !user.IsValidLenPassword() {
		return nil, &errors.ErroBase{Message: errors.UserPasswordLenErrorMessage, Code: errors.UserPasswordLenErrorCode}
	}

	if !user.IsValidUpperLetterPassword() {
		return nil, &errors.ErroBase{Message: errors.UserPasswordUpperLetterErrorMessage, Code: errors.UserPasswordUpperLetterErrorCode}
	}

	if !user.IsValidName() {
		return nil, &errors.ErroBase{Message: errors.UserNameInvalidErrorMessage, Code: errors.UserNameInvalidErrorCode}
	}

	return user, nil

}
