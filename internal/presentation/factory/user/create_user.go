package user_factories

import (
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	repository "github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
	"github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
)

func CreateUserFactoryController() controller.CreateUserController {
	service1 := service.CepServiceViaCep{}
	service2 := service.HashService{}
	repo := repository.UserRepository{}

	usecase := usecases.CreateUserUseCase{
		CepService:     service1,
		UserRepository: repo,
		HashService:    service2,
	}
	controller := controller.CreateUserController{Usecase: usecase}
	return controller
}

func CreateUserFactoryUseCase() usecases_interfaces.CreateUserUseCaseInterface {
	service1 := service.CepServiceViaCep{}
	service2 := service.HashService{}
	repo := repository.UserRepository{}

	usecase := usecases.CreateUserUseCase{
		CepService:     service1,
		UserRepository: repo,
		HashService:    service2,
	}

	return usecase
}
