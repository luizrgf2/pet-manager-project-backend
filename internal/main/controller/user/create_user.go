package controller_user

import (
	"encoding/json"
	"net/http"

	controller_error "github.com/luizrgf2/pet-manager-project-backend/internal/main/error"
	controller_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	factories "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/factory/user"
)

func CreateUserHttpController(w http.ResponseWriter, r *http.Request) {
	input := controller_interfaces.InputCreateUserController{}

	errJson := json.NewDecoder(r.Body).Decode(&input)

	w.Header().Set("Content-Type", "application/json")

	if errJson != nil {
		fieldError := controller_error.InputFieldsErrorHTTP()
		w.WriteHeader(int(fieldError.Code))
		responseJson, _ := json.Marshal(&fieldError)
		w.Write(responseJson)
		return
	}

	controller := factories.CreateUserFactoryController()
	res := controller.Handle(input)
	resJson, _ := json.Marshal(res)
	w.WriteHeader(int(res.Code))
	w.Write(resJson)
}
