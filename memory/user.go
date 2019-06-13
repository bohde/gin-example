package memory

import (
	"context"

	"github.com/joshbohde/example"
)

type UserService struct {
	Users map[int]example.User
	example.AddressService
}

func (u *UserService) User(ctx context.Context, id int) (*example.User, error) {
	user, ok := u.Users[id]
	if !ok {
		return nil, example.NotFound{}
	}

	address, err := u.AddressForUserID(ctx, id)
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
