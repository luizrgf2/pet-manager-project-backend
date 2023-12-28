package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/luizrgf2/pet-manager-project-backend/config"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	routes "github.com/luizrgf2/pet-manager-project-backend/internal/main"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	user_factories "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/factory/user"
	"github.com/stretchr/testify/assert"
)

func createUserToTest() {

	userToTeste := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        "Felicidog pet salon",
		Email:          config.SMTP_EMAIL_RECEIVER_TO_TEST,
		Password:       "Teste12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     622,
	}

	createUserUseCase := user_factories.CreateUserFactoryUseCase()

	_, err := createUserUseCase.Exec(userToTeste)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("user test 1 created!!")
	}

}

func TestSendConfirmationEmail(t *testing.T) {

	createUserToTest()

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
