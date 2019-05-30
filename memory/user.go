package memory

import (
	"github.com/joshbohde/example"
)

type UserService struct {
	Users          map[int]example.User
	AddressService example.AddressService
}

func (u *UserService) User(id int) (*example.User, error) {
	user, ok := u.Users[id]
	if !ok {
		return nil, example.NotFound{}
	}

	address, err := u.AddressService.AddressForUserId(id)
	if err != nil {
		switch err.(type) {
		case example.NotFound:
			return &user, nil
		default:
			return nil, err
		}

	}

	user.Address = address
	return &user, nil
}
