package services

type HashServiceInterface interface {
	Hash(data string) (*string, error)
	Compare(data string, encrypedData string) (bool, error)
}
