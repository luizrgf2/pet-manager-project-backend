package services

type SMTPServiceInterface interface {
	SendConfirmationEmailToUser(tokenOfConfirmation string, emailTo string) error
}
