package usecases

type InputSendConfirmationEmailToSendUserUseCase struct {
	IdUserToCreateToken uint
}

type SendConfirmationEmailToSendUserUseCase interface {
	Exec(input InputSendConfirmationEmailToSendUserUseCase) error
}
