package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	service_errors "github.com/luizrgf2/pet-manager-project-backend/internal/infra/error/services"

	services "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type responseSucess struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

type CepServiceViaCep struct {
}

func (c CepServiceViaCep) GetAddr(cep string) (*services.AddrProps, error) {
	uri := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	response, err := http.Get(uri)

	var successResponse responseSucess

	if err != nil {
		log.Fatalln(err.Error())

		return nil, &core_errors.ErroBase{
			Message: service_errors.CEPServiceUnavailableErrorMessage,
			Code:    service_errors.CEPServiceUnavailableErrorCode,
		}
	}

	if response.StatusCode == 200 {
		err_json := json.NewDecoder(response.Body).Decode(&successResponse)

		if err_json == nil && !successResponse.Erro {
			return &services.AddrProps{
				Street:   successResponse.Logradouro,
				Number:   0,
				District: successResponse.Bairro,
				City:     successResponse.Localidade,
				State:    successResponse.Uf,
			}, nil
		}

		return nil, &core_errors.ErroBase{
			Message: service_errors.CEPInvalidErrorMessage,
			Code:    service_errors.CEPInvalidErrorCode,
		}

	} else if response.StatusCode == 400 {
		return nil, &core_errors.ErroBase{
			Message: service_errors.CEPInvalidErrorMessage,
			Code:    service_errors.CEPInvalidErrorCode,
		}

	} else {
		return nil, &core_errors.ErroBase{
			Message: service_errors.CEPServiceUnavailableErrorMessage,
			Code:    service_errors.CEPServiceUnavailableErrorCode,
		}
	}

}
