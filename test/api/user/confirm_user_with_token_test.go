package user_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	user_usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_error "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	routes "github.com/luizrgf2/pet-manager-project-backend/internal/main"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"
	user_factories "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/factory/user"
	"github.com/stretchr/testify/assert"
)

func startConfirmUserApiTests() {
	CreateUserToTest()

	inputToConfirmEmail := user_usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: 1,
	}
	err := user_factories.SendConfirmationEmailToUserUseCase().Exec(inputToConfirmEmail)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestConfirmUserWithValidToken(t *testing.T) {
	startConfirmUserApiTests()

	token := FindUserConfirmationToken(1)

	requestUrl := fmt.Sprintf("/user/confirm/%s", token)

	req := httptest.NewRequest("GET", requestUrl, nil)

	r := routes.Router
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestReturnErrorIfConfirmUserAlrearyConfirmed(t *testing.T) {

	expectedError := core_errors.ErroBase{
		Message: data_error.UserAlreadyConfirmedErrorMessage,
		Code:    data_error.UserAlreadyConfirmedErrorCode,
	}

	token := FindUserConfirmationToken(1)

	requestUrl := fmt.Sprintf("/user/confirm/%s", token)

	req := httptest.NewRequest("GET", requestUrl, nil)

	r := routes.Router
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	res := contracts.HTTPResponse[interface{}]{}
	json.NewDecoder(rr.Body).Decode(&res)

	assert.Equal(t, expectedError.Code, res.Code)
	assert.Equal(t, expectedError.Message, res.ErrorsMessage[0])
}
