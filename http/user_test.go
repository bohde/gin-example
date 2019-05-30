package http

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/joshbohde/example"
	"github.com/joshbohde/example/mocks"
)

func TestUserHandler_Get(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := NewTestContext().
			SetParam("id", "1")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		users := mocks.NewMockUserService(ctrl)

		user := &example.User{ID: 1, Name: "Josh"}

		users.EXPECT().User(1).Return(user, nil)

		UserHandler{UserService: users}.Get(ctx.Context)

		ctx.AssertStatus(t, http.StatusOK)
		ctx.AssertJSONBodyEquals(t, user)
	})

	t.Run("Invalid id", func(t *testing.T) {
		ctx := NewTestContext().
			SetParam("id", "abc")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		users := mocks.NewMockUserService(ctrl)

		UserHandler{UserService: users}.Get(ctx.Context)

		ctx.AssertStatus(t, http.StatusNotFound)
	})

	t.Run("Missing user", func(t *testing.T) {
		ctx := NewTestContext().
			SetParam("id", "1")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		users := mocks.NewMockUserService(ctrl)

		users.EXPECT().User(1).Return(nil, example.NotFound{})

		UserHandler{UserService: users}.Get(ctx.Context)

		ctx.AssertStatus(t, http.StatusNotFound)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := NewTestContext().
			SetParam("id", "1")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		users := mocks.NewMockUserService(ctrl)

		users.EXPECT().User(1).Return(nil, assert.AnError)

		UserHandler{UserService: users}.Get(ctx.Context)

		ctx.AssertStatus(t, http.StatusInternalServerError)
	})
}
