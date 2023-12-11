package test

type JWTServiceInMemory struct {
}

func (J JWTServiceInMemory) CreateToken(idUser string, expirationTimeInSeconds uint) (*string, error) {
	jwtToken := "213324lkjjskdfvkjsdjeer2"
	return &jwtToken, nil
}

func (J JWTServiceInMemory) DecryptToken(token string) (*string, error) {
	password := "213324lkjjskdfvkjsdjeer2"
	return &password, nil
}
