package service_test

import (
	"testing"

	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	services_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
	infra_services_errors "github.com/luizrgf2/pet-manager-project-backend/internal/infra/error/services"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/infra/service"
	"github.com/stretchr/testify/assert"
)

var sut = services.CepServiceViaCep{}

var expectedSucessValue = services_interfaces.AddrProps{
	Street:   "Rua Alemar Rodrigues da Cunha",
	Number:   0,
	District: "Sebastião Amorim",
	City:     "Patos de Minas",
	State:    "MG",
}

func TestValidCEP(t *testing.T) {
	result, err := sut.GetAddr("38705280")
	assert.Nil(t, err)
	assert.Equal(t, &expectedSucessValue, result)
}

func TestReturnErrorWithInvalidCep(t *testing.T) {
	expctedError := &core_errors.ErroBase{
		Message: infra_services_errors.CEPInvalidErrorMessage,
		Code:    infra_services_errors.CEPInvalidErrorCode,
	}
	_, err := sut.GetAddr("02364780608")
	assert.Equal(t, expctedError, err)
}

func TestReturnErrorWithAnotherInvalidCep(t *testing.T) {
	expctedError := &core_errors.ErroBase{
		Message: infra_services_errors.CEPInvalidErrorMessage,
		Code:    infra_services_errors.CEPInvalidErrorCode,
	}
	_, err := sut.GetAddr("98705280")
	assert.Equal(t, expctedError, err)
}
