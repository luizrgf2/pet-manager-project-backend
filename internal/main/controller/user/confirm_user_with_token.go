package controller_user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	controller_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/controllers/user"
	factories "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/factory/user"
)

func ConfirmUserWithTokenHttpController(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	w.Header().Set("Content-Type", "application/json")

	input := controller_interfaces.InputConfirmUserWithTokenController{Token: token}

	controller := factories.ConfirmUserFactoryController()
	res := controller.Handle(input)
	resJson, _ := json.Marshal(res)
	w.WriteHeader(int(res.Code))
	w.Write(resJson)
}
