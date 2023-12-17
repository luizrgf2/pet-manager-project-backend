package service

import (
	"log"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

type HashService struct {
}

func (h HashService) Hash(data string) (*string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err.Error())
		return nil, &errors.ErroBase{
			Message: "Erro para codificar a senha!",
			Code:    500,
		}
	}

	hashToString := string(hashed)

	return &hashToString, nil
}

func (h HashService) Compare(data string, encrypedData string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encrypedData), []byte(data))
	return (err == nil), nil
}
