package services

type SMTPServiceInterface interface {
	SendConfirmationEmailToUser(tokenOfConfirmation string) error
}
