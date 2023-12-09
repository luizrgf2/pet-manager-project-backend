package services

type AddrProps struct {
	Street   string
	Number   uint
	District string
	City     string
	State    string
}

type CEPServiceInterface interface {
	GetAddr(cep string) (AddrProps, error)
}
