package test

import (
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	infra_error "github.com/luizrgf2/pet-manager-project-backend/internal/infra/error/services"
)

type JWTServiceInMemory struct {
}

func (J JWTServiceInMemory) CreateToken(idUser string, expirationTimeInSeconds *uint) (*string, error) {
	jwtToken := "213324lkjjskdfvkjsdjeer2"
	return &jwtToken, nil
}

func (J JWTServiceInMemory) DecryptToken(token string) (*string, error) {

	if token == "validtoken" {
		id := "1"
		return &id, nil
	} else if token == "validtoken2" {
		id := "2"
		return &id, nil
	} else {
		return nil, &core_errors.ErroBase{
			Message: infra_error.JWTIvalidTokenErrorMessage,
			Code:    infra_error.JWTIvalidTokenErrorCode,
		}
	}

}
