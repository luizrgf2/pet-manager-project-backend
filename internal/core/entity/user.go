package entity

import (
	"regexp"
	"time"
	"unicode/utf8"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
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
	AddrState      string
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

func (U *UserEntity) IsValidState() bool {
	states := []string{
		"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA",
		"MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN",
		"RS", "RO", "RR", "SC", "SP", "SE", "TO",
	}

	for _, state := range states {
		if state == U.AddrState {
			return true
		}
	}
	return false
}

func NewUser(NamePet string, Email string, Password string, AddrCep string, AddrStreet string, AddrCity string, AddrComplement string, AddrDistrict string, AddrNumber uint, AddrState string) (*UserEntity, error) {
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
		AddrState:      AddrState,
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}
	if !user.IsValidEmail() {
		return nil, &core_errors.ErroBase{Message: core_errors.UserEmailInvalidErrorMessage, Code: core_errors.UserEmailInvalidErrorCode}
	}

	if !user.IsValidLenPassword() {
		return nil, &core_errors.ErroBase{Message: core_errors.UserPasswordLenErrorMessage, Code: core_errors.UserPasswordLenErrorCode}
	}

	if !user.IsValidUpperLetterPassword() {
		return nil, &core_errors.ErroBase{Message: core_errors.UserPasswordUpperLetterErrorMessage, Code: core_errors.UserPasswordUpperLetterErrorCode}
	}

	if !user.IsValidName() {
		return nil, &core_errors.ErroBase{Message: core_errors.UserNameInvalidErrorMessage, Code: core_errors.UserNameInvalidErrorCode}
	}

	if !user.IsValidState() {
		return nil, &core_errors.ErroBase{
			Message: core_errors.UserStateInvalidErrorMessage,
			Code:    core_errors.UserStateInvalidErrorCode,
		}
	}

	return user, nil

}
