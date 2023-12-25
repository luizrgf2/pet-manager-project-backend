package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/luizrgf2/pet-manager-project-backend/config"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/data/usecase/user"
	"github.com/luizrgf2/pet-manager-project-backend/internal/infra/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	routes "github.com/luizrgf2/pet-manager-project-backend/internal/main"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	"github.com/stretchr/testify/assert"
)

func init() {
	cepService := services.CepServiceViaCep{}
	hashService := services.HashService{}
	userRepo := repository.UserRepository{}

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          config.SMTP_EMAIL_RECEIVER_TO_TEST,
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	createUserUseCase := usecases.CreateUserUseCase{
		CepService:     cepService,
		UserRepository: userRepo,
		HashService:    hashService,
	}

	_, err := createUserUseCase.Exec(userToTeste)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("user test 1 created!!")
	}

}

func TestSendConfirmationEmail(t *testing.T) {
	input := controller.InputSendEmailConfirmationToUserController{
		Id: 1,
	}
	requestBody, _ := json.Marshal(input)
	req := httptest.NewRequest("GET", "/user/confirmationemail/1", bytes.NewBuffer(requestBody))

	r := routes.Router

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
}
