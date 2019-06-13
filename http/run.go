package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshbohde/example"
)

type Server struct {
	example.AddressService
	example.UserService
}

func (s *Server) Run(port int) {
	engine := gin.Default()

	users := UserHandler{UserService: s}

	engine.GET("/users/:id", users.Get)

	engine.Run(":" + strconv.Itoa(port))

}
