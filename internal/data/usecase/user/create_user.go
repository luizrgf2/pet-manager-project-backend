package usecases

import (
	"regexp"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_errors "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type CreateUserUseCase struct {
	CepService services.CEPServiceInterface
}

func (c *CreateUserUseCase) getAddrWithCep(cep string) (*services.AddrProps, error) {
	addr, err := c.CepService.GetAddr(cep)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (c *CreateUserUseCase) validateCep(cep string) error {
	regexCEP := regexp.MustCompile(`^\d{8}$`)
	isValid := regexCEP.MatchString(cep)

	if !isValid {
		return &core_errors.ErroBase{
			Message: data_errors.CEPInvalidErrorMessage,
			Code:    uint(data_errors.CEPInvalidErrorCode),
		}
	}
	return nil
}

func (c *CreateUserUseCase) Exec(input usecases.InputCreateUserUseCase) (*usecases.OutputCreateuserUseCase, error) {

	err := c.validateCep(input.AddrCep)

	if err != nil {
		return nil, err
	}

	addr, err := c.getAddrWithCep(input.AddrCep)

	if err != nil {
		return nil, err
	}

	userData, err := entity.NewUser(
		input.NamePet,
		input.Email,
		input.Password,
		input.AddrCep,
		addr.Street,
		addr.City,
		input.AddrComplement,
		addr.District,
		input.AddrNumber,
		addr.State,
	)

	if err != nil {
		return nil, err
	}

	outputToReturn := usecases.OutputCreateuserUseCase{Id: userData.Id}

	return &outputToReturn, nil

}
