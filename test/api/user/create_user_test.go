package user_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usercases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	routes "github.com/luizrgf2/pet-manager-project-backend/internal/main"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidUser(t *testing.T) {

	input := controller.InputCreateUserController{
		NamePet:        "Felicidog Pet Salon",
		Email:          "email@email.com",
		Password:       "Test12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     786,
	}
	requestBody, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/user/create", bytes.NewBuffer(requestBody))

	r := routes.Router

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 201)

	responseStruct := contracts.HTTPResponse[usercases_interfaces.OutputCreateuserUseCase]{}

	err := json.NewDecoder(rr.Body).Decode(&responseStruct)

	assert.Nil(t, err)
	assert.Equal(t, 1, int(responseStruct.Response.Id))

}

func TestReturnErrorIfCreateUserAlreadyExists(t *testing.T) {

	expectedError := core_errors.ErroBase{
		Message: core_errors.UserAlreadyExistsErrorMessage,
		Code:    core_errors.UserAlreadyExistsErrorCode,
	}

	responseStruct := contracts.HTTPResponse[usercases_interfaces.OutputCreateuserUseCase]{}

	input := controller.InputCreateUserController{
		NamePet:        "Felicidog Pet Salon",
		Email:          "email@email.com",
		Password:       "Test12345",
		AddrCep:        "38705280",
		AddrComplement: "",
		AddrNumber:     786,
	}

	requestBody, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/user/create", bytes.NewBuffer(requestBody))

	r := routes.Router

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	err := json.NewDecoder(rr.Body).Decode(&responseStruct)

	assert.Nil(t, err)
	assert.Equal(t, rr.Code, int(expectedError.Code))
	assert.Equal(t, []string{expectedError.Message}, responseStruct.ErrorsMessage)

}
