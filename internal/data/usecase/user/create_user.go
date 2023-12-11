package usecases

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	usecases "github.com/luizrgf2/pet-manager-project-backend/internal/core/usecase/user"
	data_errors "github.com/luizrgf2/pet-manager-project-backend/internal/data/error"
	repository "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
	services "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/service"
)

type CreateUserUseCase struct {
	CepService     services.CEPServiceInterface
	UserRepository repository.UserRepositoryInterface
	HashService    services.HashServiceInterface
	JWTService     services.JWTServiceInterface
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

func (c *CreateUserUseCase) checkIfUserAlreadyExists(email string) error {
	user, err := c.UserRepository.FindByEmail(email)

	errorUserNotExists := &core_errors.ErroBase{
		Message: core_errors.UserNotExistsErrorMessage,
		Code:    core_errors.UserNotExistsErrorCode,
	}

	if errors.As(err, &errorUserNotExists) {
		return nil
	} else if err == nil && user != nil {
		return &core_errors.ErroBase{
			Message: core_errors.UserAlreadyExistsErrorMessage,
			Code:    core_errors.UserAlreadyExistsErrorCode,
		}
	} else {
		return err
	}
}

func (c *CreateUserUseCase) createUser(user *entity.UserEntity) (*uint, error) {
	user, err := c.UserRepository.Create(repository.CreateUserRepositoryInput{
		NamePet:        user.NamePet,
		Email:          user.Email,
		Password:       user.Password,
		AddrCep:        user.AddrCep,
		AddrStreet:     user.AddrStreet,
		AddrNumber:     user.AddrNumber,
		AddrComplement: &user.AddrComplement,
		AddrDistrict:   user.AddrDistrict,
		AddrCity:       user.AddrCity,
		AddrState:      user.AddrState,
	})

	if err != nil {
		return nil, err
	}
	return &user.Id, nil
}

func (c *CreateUserUseCase) saveUser(user *entity.UserEntity) (*uint, error) {
	user, err := c.UserRepository.Create(repository.CreateUserRepositoryInput{
		NamePet:        user.NamePet,
		Email:          user.Email,
		Password:       user.Password,
		AddrCep:        user.AddrCep,
		AddrStreet:     user.AddrStreet,
		AddrNumber:     user.AddrNumber,
		AddrComplement: &user.AddrComplement,
		AddrDistrict:   user.AddrDistrict,
		AddrCity:       user.AddrCity,
		AddrState:      user.AddrState,
	})
	if err != nil {
		return nil, err
	}
	return &user.Id, nil
}

func (c *CreateUserUseCase) createConfirmationToken(idUser uint) (*string, error) {
	experationTime := 86400
	token, err := c.JWTService.CreateToken(strconv.Itoa(int(idUser)), uint(experationTime))
	if err != nil {
		return nil, err
	}
	return token, nil
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

	userNotExistsOrError := c.checkIfUserAlreadyExists(input.Email)
	if userNotExistsOrError != nil {
		return nil, userNotExistsOrError
	}

	passTohash, err := c.HashService.Hash(userData.Password)
	if err != nil {
		return nil, err
	}
	userData.Password = *passTohash

	idUser, err := c.saveUser(userData)
	if err != nil {
		return nil, err
	}
	userData.Password = ""
	userData.Id = *idUser

	token, err := c.createConfirmationToken(userData.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println(token)

	outputToReturn := usecases.OutputCreateuserUseCase{Id: userData.Id}

	return &outputToReturn, nil
}
