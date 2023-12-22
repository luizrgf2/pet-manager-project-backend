package contracts

type HTTPResponse[T any] struct {
	Response      *T
	ErrorsMessage []string
	Code          uint
}
