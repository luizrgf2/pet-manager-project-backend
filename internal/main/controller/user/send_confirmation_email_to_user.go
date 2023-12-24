package controller_user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	controller_error "github.com/luizrgf2/pet-manager-project-backend/internal/main/error"
	controller_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	factories "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/factory/user"
)

func SendConfirmationEmailToUserHttpController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	w.Header().Set("Content-Type", "application/json")

	if err != nil || id < 1 {
		errorId := controller_error.InputIdInvalid()
		resJson, _ := json.Marshal(errorId)
		w.WriteHeader(int(errorId.Code))
		w.Write(resJson)
	}

	input := controller_interfaces.InputSendEmailConfirmationToUserController{Id: uint(id)}

	controller := factories.SendConfirmationEmailToUserFactoryController()
	res := controller.Handle(input)
	resJson, _ := json.Marshal(res)
	w.WriteHeader(int(res.Code))
	w.Write(resJson)
}
