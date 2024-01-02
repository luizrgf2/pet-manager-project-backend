package user_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	routes "github.com/luizrgf2/pet-manager-project-backend/internal/main"
	controller "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	"github.com/stretchr/testify/assert"
)

func TestSendConfirmationEmail(t *testing.T) {

	CreateUserToTest()

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
