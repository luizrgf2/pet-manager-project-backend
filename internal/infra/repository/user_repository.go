package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/luizrgf2/pet-manager-project-backend/internal/core/entity"
	core_errors "github.com/luizrgf2/pet-manager-project-backend/internal/core/errors"
	repository_interfaces "github.com/luizrgf2/pet-manager-project-backend/internal/data/interfaces/repository"
	DB "github.com/luizrgf2/pet-manager-project-backend/internal/infra/db"
)

type UserRepository struct {
}

func (u *UserRepository) convertDbUserToEntityUser(entityUser *entity.UserEntity, dbUser *sql.Rows) {
	dbUser.Scan(
		&entityUser.Id,
		&entityUser.NamePet,
		&entityUser.Email,
		&entityUser.Password,
		&entityUser.AddrCep,
		&entityUser.AddrStreet,
		&entityUser.AddrNumber,
		&entityUser.AddrDistrict,
		&entityUser.AddrCity,
		&entityUser.AddrState,
		&entityUser.AddrComplement,
		nil,
		nil,
		&entityUser.CreatedAt,
		&entityUser.UpdatedAt,
	)
}

func (u UserRepository) Create(input repository_interfaces.CreateUserRepositoryInput) (*entity.UserEntity, error) {

	createdAt := time.Now()
	updatedAt := time.Now()

	insertStatement := `
		INSERT INTO users (name, email, password, addr_cep, addr_street, addr_number, addr_district, addr_city, addr_state, addr_complement, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := DB.DB.Exec(
		insertStatement,
		input.NamePet,
		input.Email,
		input.Password,
		input.AddrCep,
		input.AddrStreet,
		input.AddrNumber,
		input.AddrDistrict,
		input.AddrCity,
		input.AddrState,
		input.AddrComplement,
		createdAt,
		updatedAt,
	)

	if err != nil {
		return nil, &core_errors.ErroBase{Message: "Erro para criar o usuário!", Code: 500}
	}

	idInserted, err := result.LastInsertId()

	if err != nil {
		return nil, &core_errors.ErroBase{Message: "Erro para criar o usuário!", Code: 500}
	}

	return &entity.UserEntity{
		Id:             uint(idInserted),
		NamePet:        input.NamePet,
		Email:          input.Email,
		Password:       "",
		AddrCep:        input.AddrCep,
		AddrStreet:     input.AddrStreet,
		AddrNumber:     input.AddrNumber,
		AddrComplement: *input.AddrComplement,
		AddrDistrict:   input.AddrDistrict,
		AddrCity:       input.AddrCity,
		AddrState:      input.AddrState,
		UpdatedAt:      createdAt,
		CreatedAt:      updatedAt,
	}, nil
}

func (u UserRepository) FindById(id uint) (*entity.UserEntity, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = %d", id)

	result, err := DB.DB.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	user := entity.UserEntity{}

	for result.Next() {
		u.convertDbUserToEntityUser(&user, result)
	}

	if user.Id == 0 {
		return nil, &core_errors.ErroBase{
			Message: core_errors.UserNotExistsErrorMessage,
			Code:    core_errors.UserNotExistsErrorCode,
		}
	}

	return &user, nil
}

func (u UserRepository) FindByEmail(email string) (*entity.UserEntity, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)

	result, err := DB.DB.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	user := entity.UserEntity{}

	for result.Next() {
		u.convertDbUserToEntityUser(&user, result)
	}

	if user.Id == 0 {
		return nil, &core_errors.ErroBase{
			Message: core_errors.UserNotExistsErrorMessage,
			Code:    core_errors.UserNotExistsErrorCode,
		}
	}

	return &user, nil
}

func (u UserRepository) UpdateConfirmationToken(id uint, token string, expirationTimeInSeconds *uint) error {

	if expirationTimeInSeconds != nil {
		expiration := *expirationTimeInSeconds
		expirationDate := time.Now()
		expirationDate.Add(time.Duration(expiration))

		query := fmt.Sprintf("UPDATE users SET confirmation_token = '%s', expiration_confirmation_token='%s' WHERE id=%d", token, expirationDate.Format("2006-01-02 15:04:05"), id)
		result, err := DB.DB.Exec(query)

		if err != nil {
			return &core_errors.ErroBase{
				Message: "Erro para atualizar o token de confirmação!",
				Code:    500,
			}
		}

		if rowsEffected, err := result.RowsAffected(); err != nil {
			if rowsEffected != 1 {
				return &core_errors.ErroBase{
					Message: "Erro para atualizar o token de confirmação!",
					Code:    500,
				}
			}
		} else {
			return &core_errors.ErroBase{
				Message: "Erro para atualizar o token de confirmação!",
				Code:    500,
			}
		}

	} else {
		query := fmt.Sprintf("UPDATE users SET confirmation_token = '%s' WHERE id=%d", token, id)
		result, err := DB.DB.Exec(query)

		if err != nil {
			return &core_errors.ErroBase{
				Message: "Erro para atualizar o token de confirmação!",
				Code:    500,
			}
		}

		if rowsEffected, err := result.RowsAffected(); err != nil {
			if rowsEffected != 1 {
				return &core_errors.ErroBase{
					Message: "Erro para atualizar o token de confirmação!",
					Code:    500,
				}
			}
		} else {
			return &core_errors.ErroBase{
				Message: "Erro para atualizar o token de confirmação!",
				Code:    500,
			}
		}
	}
	return nil
}

func (u UserRepository) Update(id uint, input repository_interfaces.UpdateUserRepositoryInput) (*entity.UserEntity, error) {
	return &entity.UserEntity{
		Id:             0,
		NamePet:        "ValidName",
		Email:          "validemail@gmail.com",
		Password:       "ValidPass123",
		AddrCep:        "38705280",
		AddrStreet:     "Alemar Rodrigues da cunha",
		AddrNumber:     622,
		AddrComplement: "",
		AddrDistrict:   "Sebastião Amorim",
		AddrCity:       "Patos de Minas",
		AddrState:      "MG",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}, nil
}

func (u UserRepository) Delete(id uint) error {
	return nil
}
