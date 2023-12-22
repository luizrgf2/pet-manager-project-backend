package contracts

type HTTPResponse[T any] struct {
	Response      *T       `json:"response"`
	ErrorsMessage []string `json:"errors"`
	Code          uint     `json:"code"`
}
