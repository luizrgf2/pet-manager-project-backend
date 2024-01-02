package user_test

import (
	"fmt"
	"log"

	"github.com/luizrgf2/pet-manager-project-backend/config"
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	user_usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	repository_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
	"github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
)

func CreateUserToTest() {
	userToTeste := user_usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          config.SMTP_EMAIL_RECEIVER_TO_TEST,
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	input := repository_interfaces.CreateUserRepositoryInput{
		NamePet:        userToTeste.NamePet,
		Email:          userToTeste.Email,
		Password:       userToTeste.Password,
		AddrCep:        userToTeste.AddrCep,
		AddrStreet:     "Alemar Rodrigues da Cunha",
		AddrNumber:     622,
		AddrComplement: &userToTeste.AddrComplement,
		AddrDistrict:   "Sebastião Amorim",
		AddrCity:       "Patos de Minas",
		AddrState:      "MG",
	}

	userRepo := repository.UserRepository{}

	_, err := userRepo.Create(input)
	if err != nil {
		log.Fatalf("Erro para criar o usuário!")
	} else {
		fmt.Println("Sucesso para criar o usuário!")
	}
}

func FindUserToTest(idUser uint) entity.UserEntity {
	userRepo := repository.UserRepository{}
	res, err := userRepo.FindById(idUser)
	if err != nil {
		log.Fatalf("Erro para pegar o token de confirmação!")
	}
	return *res
}

func FindUserConfirmationToken(idUser uint) string {
	userRepo := repository.UserRepository{}
	res, err := userRepo.FindConfirmationToken(idUser)
	if err != nil {
		log.Fatalf("Erro para pegar o token de confirmação!")
	}
	return *res
}
