package usecases

type InputSendConfirmationEmailToSendUserUseCase struct {
	idUserToCreateToken string
}

type SendConfirmationEmailToSendUserUseCase interface {
	Exec(input InputSendConfirmationEmailToSendUserUseCase) error
}
