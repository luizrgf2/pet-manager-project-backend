package test

import (
	erros_core "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	erros_data "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"

	interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type CEPServiceInMemory struct {
}

func (c *CEPServiceInMemory) cepsForTest(cep string) (*interfaces.AddrProps, error) {
	cepToTeste := interfaces.AddrProps{
		Street:   "Alemar Rodrigues Da Cunha",
		Number:   622,
		District: "Sebastião Amotim",
		City:     "Patos de Minas",
		State:    "MG",
	}

	if cep == "38705280" {
		return &cepToTeste, nil
	} else {
		return nil, &erros_core.ErroBase{Message: erros_data.CEPNotExistsErrorMessage, Code: uint(erros_data.CEPNotExistsErrorCode)}
	}

}

func (c CEPServiceInMemory) GetAddr(cep string) (*interfaces.AddrProps, error) {
	addr, err := c.cepsForTest(cep)

	if err != nil {
		return nil, err
	}

	return addr, nil

}
