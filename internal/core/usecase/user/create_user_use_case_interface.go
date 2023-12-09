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
	id string
}

type CreateUserUseCaseInterface interface {
	exec(input InputCreateUserUseCase) (OutputCreateuserUseCase, error)
}
