package user_controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	usecases_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/common"
	"github.com/luizrgf2/pet-manager-project-backend/internal/presentation/contracts"
)

type InputSendEmailConfirmationToUserController struct {
	Id uint `json:"id" validator:"required,number,min=1"`
}

type SendConfirmationEmailToUserController struct {
	Usecase usecases_interfaces.SendConfirmationEmailToSendUserUseCase
}

func (c *SendConfirmationEmailToUserController) validateFields(input InputSendEmailConfirmationToUserController) []string {
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

func (c *SendConfirmationEmailToUserController) Handle(input InputSendEmailConfirmationToUserController) contracts.HTTPResponse[usecases_interfaces.OutputCreateuserUseCase] {

	errorHandling := common.ErrorHandling[usecases_interfaces.OutputCreateuserUseCase]{}

	errorsToValidateFields := c.validateFields(input)
	if len(errorsToValidateFields) > 0 {
		return errorHandling.HandlingFieldicError(errorsToValidateFields)
	}

	inputUseCase := usecases_interfaces.InputSendConfirmationEmailToSendUserUseCase{
		IdUserToCreateToken: input.Id,
	}

	err := c.Usecase.Exec(inputUseCase)

	if err != nil {
		return errorHandling.HandlingGenericError(err)
	}

	return contracts.HTTPResponse[usecases_interfaces.OutputCreateuserUseCase]{
		Response:      nil,
		ErrorsMessage: []string{},
		Code:          200,
	}
}
