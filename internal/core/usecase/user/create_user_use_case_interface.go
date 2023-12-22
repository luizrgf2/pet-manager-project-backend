package usecases

type InputCreateUserUseCase struct {
	NamePet        string
	Email          string
	Password       string
	AddrCep        string
	AddrComplement string
	AddrNumber     uint
}

type OutputCreateuserUseCase struct {
	Id uint `json:"id"`
}

type CreateUserUseCaseInterface interface {
	Exec(input InputCreateUserUseCase) (*OutputCreateuserUseCase, error)
}
