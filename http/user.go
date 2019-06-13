package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/joshbohde/example"
)

type UserHandler struct {
	example.UserService
}

func (h UserHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		// Params that aren't ints return 404
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user, err := h.User(c.Request.Context(), id)

	if err != nil {
		switch err.(type) {
		case example.NotFound:
			c.AbortWithStatus(http.StatusNotFound)
		default:
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
