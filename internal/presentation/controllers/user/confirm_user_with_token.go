package user_controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/common"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"
)

type InputConfirmUserWithTokenController struct {
	Token string `json:"token" validator:"required"`
}

type ConfirmUserWithTokenController struct {
	Usecase usecases_interfaces.CofirmUserWithTokenUseCaseInterface
}

func (c *ConfirmUserWithTokenController) validateFields(input InputConfirmUserWithTokenController) []string {
	errorsToReturn := []string{}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err_validations := validate.Struct(&input)
	if err_validations == nil {
		return errorsToReturn
	}
	for _, e := range err_validations.(validator.ValidationErrors) {
		if e != nil {
			errorsToReturn = append(errorsToReturn, fmt.Sprintf("Erro no campo [%s] : %s", e.Field(), e.Error()))
		}
	}
	return errorsToReturn
}

func (c *ConfirmUserWithTokenController) Handle(input InputConfirmUserWithTokenController) contracts.HTTPResponse[usecases_interfaces.OutputConfirmUserWithTokenUseCase] {

	errorHandling := common.ErrorHandling[usecases_interfaces.OutputConfirmUserWithTokenUseCase]{}

	errorsToValidateFields := c.validateFields(input)
	if len(errorsToValidateFields) > 0 {
		return errorHandling.HandlingFieldicError(errorsToValidateFields)
	}

	input_usecase := usecases_interfaces.InputConfirmUserWithTokenUseCase{
		Token: input.Token,
	}

	res, err := c.Usecase.Exec(input_usecase)

	if err != nil {
		return errorHandling.HandlingGenericError(err)
	}

	outputRes := usecases_interfaces.OutputConfirmUserWithTokenUseCase{
		NamePet: res.NamePet,
	}

	return contracts.HTTPResponse[usecases_interfaces.OutputConfirmUserWithTokenUseCase]{
		Response:      &outputRes,
		ErrorsMessage: []string{},
		Code:          200,
	}
}
