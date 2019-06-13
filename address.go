package example

import "context"

//go:generate mockgen -source=address.go -package=mocks -destination=mocks/address.go

type Address struct {
	ID      int
	Street  string
	ZipCode string
}

type AddressService interface {
	AddressForUserId(ctx context.Context, id int) (*Address, error)
}
