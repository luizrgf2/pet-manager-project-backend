package test

type SMTPServiceInMemory struct {
}

func (s SMTPServiceInMemory) SendConfirmationEmailToUser(tokenOfConfirmation string) error {
	return nil
}
