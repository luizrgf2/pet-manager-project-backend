package controller_error

import "github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"

func InputFieldsErrorHTTP() contracts.HTTPResponse[interface{}] {
	return contracts.HTTPResponse[interface{}]{
		Response:      nil,
		ErrorsMessage: []string{"Erro nos campos de entrada!"},
		Code:          400,
	}
}
