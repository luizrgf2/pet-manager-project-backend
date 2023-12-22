package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/common"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"
)

type InputCreateUserController struct {
	NamePet        string `validate:"required"`
	Email          string `validate:"required, email"`
	Password       string `validate:"required, max=30, min=8"`
	AddrCep        string `validate:"required, max=8, min=8"`
	AddrComplement string `validate:"max=100"`
	AddrNumber     uint   `validate:"required"`
}

type CreateUserController struct {
	Usecase usecases_interfaces.CreateUserUseCaseInterface
}

func (c *CreateUserController) validateFields(input InputCreateUserController) []string {
	errorsToReturn := []string{}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err_validations := validate.Struct(input)
	for _, e := range err_validations.(validator.ValidationErrors) {
		errorsToReturn = append(errorsToReturn, fmt.Sprintf("Erro no campo [%s] : %s", e.Field(), e.Error()))
	}
	return errorsToReturn
}

func (c *CreateUserController) Handle(input InputCreateUserController) contracts.HTTPResponse[usecases_interfaces.OutputCreateuserUseCase] {

	errorHandling := common.ErrorHandling[usecases_interfaces.OutputCreateuserUseCase]{}

	errorsToValidateFields := c.validateFields(input)
	if len(errorsToValidateFields) > 0 {
		return errorHandling.HandlingFieldicError(errorsToValidateFields)
	}

	input_usecase := usecases_interfaces.InputCreateUserUseCase{
		NamePet:        input.NamePet,
		Email:          input.Email,
		Password:       input.Password,
		AddrCep:        input.AddrCep,
		AddrComplement: input.AddrComplement,
		AddrNumber:     input.AddrNumber,
	}

	res, err := c.Usecase.Exec(input_usecase)

	if err != nil {
		return errorHandling.HandlingGenericError(err)
	}

	outputRes := usecases_interfaces.OutputCreateuserUseCase{
		Id: res.Id,
	}

	return contracts.HTTPResponse[usecases_interfaces.OutputCreateuserUseCase]{
		Response:      &outputRes,
		ErrorsMessage: []string{},
		Code:          201,
	}
}
