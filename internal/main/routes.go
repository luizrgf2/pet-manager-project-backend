package routes

import (
	"github.com/gorilla/mux"
	controller_user "github.com/luizrgf2/pet-manager-project-backend/internal/main/controller/user"
)

var Router = mux.NewRouter()

func userRoutes() {
	userRouters := Router.PathPrefix("/user").Subrouter()

	//public routes
	userRouters.HandleFunc("/create", controller_user.CreateUserHttpController).Methods("POST")
	userRouters.HandleFunc("/confirmationemail/{id}", controller_user.SendConfirmationEmailToUserHttpController).Methods("GET")

}

func init() {
	userRoutes()
}
