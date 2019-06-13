package memory

import (
	"context"

	"github.com/joshbohde/example"
)

type AddressService struct {
	Addresses map[int]example.Address
}

func (a *AddressService) AddressForUserId(ctx context.Context, id int) (*example.Address, error) {
	addr, ok := a.Addresses[id]
	if !ok {
		return nil, example.NotFound{}
	}

	return &addr, nil
}
