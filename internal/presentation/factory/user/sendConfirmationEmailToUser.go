package factories

import (
	"github.com/luizrgf2/pet-manager-project-backend/config"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	repository "github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
	"github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
)

func SendConfirmationEmailToUserFactoryController() controller.SendConfirmationEmailToUserController {
	service1 := service.JWTService{}
	service2 := service.SMTPService{}
	repo := repository.UserRepository{}

	usecase := usecases.SendConfirmationEmailToUserUseCase{
		JwtService:                      service1,
		SMTPService:                     service2,
		ExpirationTimeForTokenInSeconds: uint(config.EXPIRATION_TIME_IN_SECONDS),
		UserRepo:                        repo,
	}
	controller := controller.SendConfirmationEmailToUserController{Usecase: usecase}
	return controller
}

func SendConfirmationEmailToUserUseCase() usecases_interfaces.SendConfirmationEmailToSendUserUseCase {
	service1 := service.JWTService{}
	service2 := service.SMTPService{}
	repo := repository.UserRepository{}

	usecase := usecases.SendConfirmationEmailToUserUseCase{
		JwtService:                      service1,
		SMTPService:                     service2,
		ExpirationTimeForTokenInSeconds: uint(config.EXPIRATION_TIME_IN_SECONDS),
		UserRepo:                        repo,
	}

	return usecase
}
