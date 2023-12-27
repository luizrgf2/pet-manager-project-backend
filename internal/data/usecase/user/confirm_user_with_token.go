package user_usercases

import (
	"strconv"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	user_usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_error "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	"github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type ConfirmUserWithToken struct {
	JwtService services.JWTServiceInterface
	UserRepo   repository.UserRepositoryInterface
}

func (c *ConfirmUserWithToken) decodeToken(token string) (uint, error) {
	tokenDecoded, err := c.JwtService.DecryptToken(token)
	if err != nil {
		return 0, err
	}

	idConverded, err := strconv.Atoi(*tokenDecoded)
	if err != nil {
		return 0, &core_errors.ErroBase{
			Message: "Erro para converter o id do usuário!",
			Code:    500,
		}
	}
	return uint(idConverded), nil
}

func (c *ConfirmUserWithToken) findUserById(id uint) (*entity.UserEntity, error) {
	user, err := c.UserRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c ConfirmUserWithToken) Exec(input user_usecases_interfaces.InputConfirmUserWithTokenUseCase) (*user_usecases_interfaces.OutputConfirmUserWithTokenUseCase, error) {
	idUser, err := c.decodeToken(input.Token)
	if err != nil {
		return nil, err
	}

	confirmed, err := c.UserRepo.CheckIfUserConfirmed(idUser)
	if err != nil {
		return nil, err
	}
	if confirmed {
		return nil, &core_errors.ErroBase{
			Message: data_error.UserAlreadyConfirmedErrorMessage,
			Code:    data_error.UserAlreadyConfirmedErrorCode,
		}
	}

	user, err := c.findUserById(idUser)
	if err != nil {
		return nil, err
	}

	return &user_usecases_interfaces.OutputConfirmUserWithTokenUseCase{NamePet: user.NamePet}, nil

}
