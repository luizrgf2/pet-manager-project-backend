package user_usercases_test

import (
	"testing"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	user_usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_error "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	user_usercases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	infra_error "github.com/luizrgf2/pet-manager-project-backend/internal/infra/error/services"
	repository "github.com/luizrgf2/pet-manager-project-backend/test/moks/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/test/moks/service"
	"github.com/stretchr/testify/assert"
)

func TestConfirmUserWithValidToken(t *testing.T) {

	jwtService := services.JWTServiceInMemory{}
	userRepo := repository.UserRepositoryInMemory{}

	sut := user_usercases.ConfirmUserWithToken{
		JwtService: jwtService,
		UserRepo:   userRepo,
	}

	input := user_usecases_interfaces.InputConfirmUserWithTokenUseCase{
		Token: "validtoken",
	}

	_, err := sut.Exec(input)
	assert.Nil(t, err)

}

func TestReturnErroIfConfirmUserWithInvalidToken(t *testing.T) {

	expectedError := &core_errors.ErroBase{
		Message: infra_error.JWTIvalidTokenErrorMessage,
		Code:    infra_error.JWTIvalidTokenErrorCode,
	}

	jwtService := services.JWTServiceInMemory{}
	userRepo := repository.UserRepositoryInMemory{}

	sut := user_usercases.ConfirmUserWithToken{
		JwtService: jwtService,
		UserRepo:   userRepo,
	}

	input := user_usecases_interfaces.InputConfirmUserWithTokenUseCase{
		Token: "invalidtoken2",
	}

	_, err := sut.Exec(input)
	assert.Equal(t, expectedError, err)

}

func TestReturnErroIfConfirmUserAlrearyConfirmed(t *testing.T) {

	expectedError := &core_errors.ErroBase{
		Message: data_error.UserAlreadyConfirmedErrorMessage,
		Code:    data_error.UserAlreadyConfirmedErrorCode,
	}

	jwtService := services.JWTServiceInMemory{}
	userRepo := repository.UserRepositoryInMemory{}

	sut := user_usercases.ConfirmUserWithToken{
		JwtService: jwtService,
		UserRepo:   userRepo,
	}

	input := user_usecases_interfaces.InputConfirmUserWithTokenUseCase{
		Token: "validtoken2",
	}

	_, err := sut.Exec(input)
	assert.Equal(t, expectedError, err)

}
