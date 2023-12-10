package services

type JWTService interface {
	createToken(idUser string, expirationTimeInSeconds *uint) (string, error)
	decryptToken(token string) (string, error)
}
