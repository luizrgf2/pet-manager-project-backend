package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/luizrgf2/pet-manager-project-backend/config"
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	jwt_errors "github.com/luizrgf2/pet-manager-project-backend/internal/infra/error/services"
)

type myCustomClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

type JWTService struct{}

func (j JWTService) CreateToken(idUser string, expirationTimeInSeconds *uint) (*string, error) {
	if expirationTimeInSeconds != nil {
		expiration := *expirationTimeInSeconds
		expirationDate := time.Now()
		expirationDate.Add(time.Duration(expiration))

		registerClains := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
		}

		clains := &myCustomClaims{
			Id:               idUser,
			RegisteredClaims: registerClains,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES256, clains)

		tokenString, err := token.SignedString(config.JWT_KEY)
		if err != nil {
			return nil, &core_errors.ErroBase{
				Message: jwt_errors.JWTErrorToCreateTokenErrorMessage,
				Code:    uint(jwt_errors.JWTErrorToCreateTokenErrorCode),
			}
		}

		return &tokenString, nil
	} else {

		clains := &myCustomClaims{
			Id: idUser,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
		tokenString, err := token.SignedString([]byte(config.JWT_KEY))
		if err != nil {
			return nil, &core_errors.ErroBase{
				Message: jwt_errors.JWTErrorToCreateTokenErrorMessage,
				Code:    uint(jwt_errors.JWTErrorToCreateTokenErrorCode),
			}
		}

		return &tokenString, nil
	}
}

func (j JWTService) DecryptToken(token string) (*string, error) {

	validate := func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	}

	token_res, err := jwt.ParseWithClaims(token, &myCustomClaims{}, validate)

	if err != nil {
		return nil, &core_errors.ErroBase{
			Message: jwt_errors.JWTIvalidTokenErrorMessage,
			Code:    jwt_errors.CEPInvalidErrorCode,
		}
	}

	if clains, ok := token_res.Claims.(*myCustomClaims); ok {
		return &clains.Id, nil
	} else {
		return nil, &core_errors.ErroBase{
			Message: jwt_errors.JWTIvalidTokenErrorMessage,
			Code:    jwt_errors.CEPInvalidErrorCode,
		}
	}
}
