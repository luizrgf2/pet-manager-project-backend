package errors

const (
	UserPasswordLenErrorCode         = 422
	UserPasswordUpperLetterErrorCode = 422
	UserEmailInvalidErrorCode        = 422
	UserNameInvalidErrorCode         = 422
	UserNotExistsErrorCode           = 404
	UserAlreadyExistsErrorCode       = 409
	UserStateInvalidErrorCode        = 422
)

const (
	UserPasswordLenErrorMessage         = "A senha deve ter entre 8 e 15 caracteres!"
	UserPasswordUpperLetterErrorMessage = "A senha deve ter pelo menos uma letra maiúscula!"
	UserEmailInvalidErrorMessage        = "O email é inválido!"
	UserNameInvalidErrorMessage         = "O nome deve ter de 4 a 50 caracteres!"
	UserNotExistsErrorMessage           = "O usuário não existe!"
	UserAlreadyExistsErrorMessage       = "O usuário já existe!"
	UserStateInvalidErrorMessage        = "O estado de localização não é válido!"
)
