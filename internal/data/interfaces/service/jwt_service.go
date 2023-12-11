package services

type JWTServiceInterface interface {
	CreateToken(idUser string, expirationTimeInSeconds uint) (*string, error)
	DecryptToken(token string) (*string, error)
}
