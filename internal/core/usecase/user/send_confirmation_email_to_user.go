package user_usecases_interfaces

type InputSendConfirmationEmailToSendUserUseCase struct {
	IdUserToCreateToken uint
}

type SendConfirmationEmailToSendUserUseCase interface {
	Exec(input InputSendConfirmationEmailToSendUserUseCase) error
}
