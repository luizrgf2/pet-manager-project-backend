package user_usecases_interfaces

type InputConfirmUserWithTokenUseCase struct {
	token string
}

type OutputConfirmUserWithTokenUseCase struct {
	namePet string
}

type CofirmUserWithTokenUseCaseInterface interface {
	Exec(input InputConfirmUserWithTokenUseCase) (OutputConfirmUserWithTokenUseCase, error)
}
