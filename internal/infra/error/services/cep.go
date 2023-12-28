package infra_error

var (
	CEPServiceUnavailableErrorMessage = "O serviço de cep não está disponível!"
	CEPInvalidErrorMessage            = "O cep informado é inválido!"
)

var (
	CEPServiceUnavailableErrorCode = uint(503)
	CEPInvalidErrorCode            = uint(400)
)
