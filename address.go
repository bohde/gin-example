package example

//go:generate mockgen -source=address.go -package=mocks -destination=mocks/address.go

type Address struct {
	ID      int
	Street  string
	ZipCode string
}

type AddressService interface {
	AddressForUserId(id int) (*Address, error)
}
