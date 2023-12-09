package usecases

import (
	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	"github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/services"
)

type CreateUserUseCase struct {
	cepService services.CEPServiceInterface
}

func (c *CreateUserUseCase) getAddrWithCep(cep string) (*services.AddrProps, error) {
	addr, err := c.cepService.GetAddr(cep)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func (c *CreateUserUseCase) Exec(input usecases.InputCreateUserUseCase) (*usecases.OutputCreateuserUseCase, error) {

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
	)

	if err != nil {
		return nil, err
	}

	outputToReturn := usecases.OutputCreateuserUseCase{Id: userData.Id}

	return &outputToReturn, nil

}
