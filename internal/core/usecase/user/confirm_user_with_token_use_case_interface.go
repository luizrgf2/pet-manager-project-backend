package user_usecases_interfaces

type InputConfirmUserWithTokenUseCase struct {
	Token string
}

type OutputConfirmUserWithTokenUseCase struct {
	NamePet string
}

type CofirmUserWithTokenUseCaseInterface interface {
	Exec(input InputConfirmUserWithTokenUseCase) (*OutputConfirmUserWithTokenUseCase, error)
}
