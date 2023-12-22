package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luizrgf2/pet-manager-project-backend/config"
	routes "github.com/luizrgf2/pet-manager-project-backend/internal/main"
)

func main() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), routes.Router)
	if err != nil {
		log.Fatal("Erro para iniciar o server", err.Error())
	}
	fmt.Println("Sucess sever is running")
}
