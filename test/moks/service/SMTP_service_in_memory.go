package test

type SMTPServiceInMemory struct {
}

func (s SMTPServiceInMemory) SendConfirmationEmailToUser(tokenOfConfirmation string, emailTo string) error {
	return nil
}
