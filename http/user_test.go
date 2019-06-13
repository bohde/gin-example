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
	type user struct {
		user *example.User
		err  error
	}

	type out struct {
		status int
		body   interface{}
	}

	u := &example.User{ID: 1, Name: "Josh"}

	cases := map[string]struct {
		param    string
		user     *user
		expected out
	}{
		"Success": {
			param:    "1",
			user:     &user{user: u},
			expected: out{status: http.StatusOK, body: u},
		},
		"InvalidID": {
			param:    "abc",
			expected: out{status: http.StatusNotFound},
		},
		"UserNotFound": {
			param:    "1",
			user:     &user{err: example.NotFound{}},
			expected: out{status: http.StatusNotFound},
		},
		"Error": {
			param:    "1",
			user:     &user{err: assert.AnError},
			expected: out{status: http.StatusInternalServerError},
		},
	}

	for id, c := range cases {
		c := c
		t.Run(id, func(t *testing.T) {
			ctx := NewTestContext().
				SetParam("id", c.param)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			users := mocks.NewMockUserService(ctrl)

			if c.user != nil {
				users.EXPECT().User(ctx.Context.Request.Context(), 1).Return(c.user.user, c.user.err)
			}

			UserHandler{users}.Get(ctx.Context)

			ctx.AssertStatus(t, c.expected.status)

			if c.expected.body != nil {
				ctx.AssertJSONBodyEquals(t, c.expected.body)
			}
		})

	}
}
