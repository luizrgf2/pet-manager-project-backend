package services

type SMTPService interface {
	SendConfirmationEmailToUser(tokenOfConfirmation string) error
}
