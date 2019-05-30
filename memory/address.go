package memory

import (
	"github.com/joshbohde/example"
)

type AddressService struct {
	Addresses map[int]example.Address
}

func (a *AddressService) AddressForUserId(id int) (*example.Address, error) {
	addr, ok := a.Addresses[id]
	if !ok {
		return nil, example.NotFound{}
	}

	return &addr, nil
}
