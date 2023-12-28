package user_factories

import (
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	repository "github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
	"github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
)

func ConfirmUserFactoryController() controller.ConfirmUserWithTokenController {
	service1 := service.JWTService{}
	repo := repository.UserRepository{}

	usecase := usecases.ConfirmUserWithToken{
		JwtService: service1,
		UserRepo:   repo,
	}
	controller := controller.ConfirmUserWithTokenController{Usecase: usecase}
	return controller
}

func ConfirmUserFactoryUseCase() usecases_interfaces.CofirmUserWithTokenUseCaseInterface {
	service1 := service.JWTService{}
	repo := repository.UserRepository{}

	usecase := usecases.ConfirmUserWithToken{
		JwtService: service1,
		UserRepo:   repo,
	}

	return usecase
}
