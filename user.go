package example

import "context"

//go:generate mockgen -source=user.go -package=mocks -destination=mocks/user.go

type User struct {
	ID      int
	Name    string
	Address *Address
}

type UserService interface {
	User(ctx context.Context, id int) (*User, error)
}
