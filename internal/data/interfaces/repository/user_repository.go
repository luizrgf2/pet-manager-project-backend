package repository

type CreateUserRepositoryInput struct {
	NamePet        string
	Email          string
	Password       string
	AddrStreet     string
	AddrNumber     uint
	AddrComplement *string
	AddrDistrict   string
	AddrCity       string
	AddrState      string
}

type UserRepositoryInterface interface {
	create()
}
