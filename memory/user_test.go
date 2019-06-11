package memory

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/joshbohde/example"
	"github.com/joshbohde/example/mocks"
)

func TestUserService(t *testing.T) {
	type addr struct {
		address *example.Address
		err     error
	}

	type out struct {
		user *example.User
		err  error
	}

	address := example.Address{ID: 2, Street: "Foo", ZipCode: "Baz"}

	cases := map[string]struct {
		userID   int   // User ID to search for
		address  *addr // The potentially nil addr to return
		expected out   // the expected
	}{
		"Good": {
			userID:   1,
			address:  &addr{address: &address},
			expected: out{user: &example.User{ID: 1, Name: "Josh", Address: &address}},
		},
		"UserNotFound": {
			userID:   2,
			expected: out{err: example.NotFound{}},
		},
		"AddressNotFound": {
			userID:   1,
			address:  &addr{err: example.NotFound{}},
			expected: out{user: &example.User{ID: 1, Name: "Josh"}},
		},
		"AddressErrors": {
			userID:   1,
			address:  &addr{err: assert.AnError},
			expected: out{err: assert.AnError},
		},
	}

	for id, c := range cases {

		t.Run(id, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			addresses := mocks.NewMockAddressService(ctrl)

			users := UserService{
				Users: map[int]example.User{
					1: example.User{ID: 1, Name: "Josh"},
				},
				AddressService: addresses,
			}

			if c.address != nil {
				addresses.EXPECT().AddressForUserId(c.userID).Return(c.address.address, c.address.err)
			}
			user, err := users.User(c.userID)

			assert.Equal(t, c.expected, out{user, err})
		})
	}
}
