package service_test

import (
	"testing"

	services "github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidToken(t *testing.T) {
	sut_jwt := services.JWTService{}
	result, err := sut_jwt.CreateToken("1", nil)
	assert.Nil(t, err)
	assert.NotNil(t, result)

}

func TestDecodeToken(t *testing.T) {
	sut_jwt := services.JWTService{}

	token, err := sut_jwt.CreateToken("1", nil)

	result, err := sut_jwt.DecryptToken(*token)
	assert.Nil(t, err)
	assert.Equal(t, *result, "1")
}
