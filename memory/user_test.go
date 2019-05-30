package memory

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joshbohde/example"
	"github.com/joshbohde/example/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	t.Run("User", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		addresses := mocks.NewMockAddressService(ctrl)

		users := UserService{
			users: map[int]example.User{
				1: example.User{ID: 1, Name: "Josh"},
			},
			AddressService: addresses,
		}

		address := example.Address{ID: 2, Street: "Foo", ZipCode: "Baz"}

		addresses.EXPECT().AddressForUserId(1).Return(&address, nil)

		expected := example.User{
			ID:      1,
			Name:    "Josh",
			Address: &address,
		}
		user, err := users.User(1)

		assert.Nil(t, err)
		assert.Equal(t, &expected, user)

	})

	t.Run("Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		addresses := mocks.NewMockAddressService(ctrl)

		users := UserService{
			AddressService: addresses,
		}
		user, err := users.User(1)

		assert.Equal(t, example.NotFound{}, err)
		assert.Nil(t, user)
	})

	t.Run("Address Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		addresses := mocks.NewMockAddressService(ctrl)

		users := UserService{
			users: map[int]example.User{
				1: example.User{ID: 1, Name: "Josh"},
			},
			AddressService: addresses,
		}

		addresses.EXPECT().AddressForUserId(1).Return(nil, example.NotFound{})

		expected := example.User{
			ID:   1,
			Name: "Josh",
		}
		user, err := users.User(1)

		assert.Nil(t, err)
		assert.Equal(t, &expected, user)

	})

	t.Run("Address Errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		addresses := mocks.NewMockAddressService(ctrl)

		users := UserService{
			users: map[int]example.User{
				1: example.User{ID: 1, Name: "Josh"},
			},
			AddressService: addresses,
		}

		addresses.EXPECT().AddressForUserId(1).Return(nil, assert.AnError)

		user, err := users.User(1)

		assert.Nil(t, user)
		assert.Equal(t, err, assert.AnError)

	})

}
