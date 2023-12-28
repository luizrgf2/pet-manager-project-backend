package infra_error

var (
	JWTIvalidTokenErrorMessage        = "O token fornecido não é válido!"
	JWTErrorToCreateTokenErrorMessage = "Aconteceu algum erro na criação do token!"
)

var (
	JWTIvalidTokenErrorCode        = uint(401)
	JWTErrorToCreateTokenErrorCode = 500
)
