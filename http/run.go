package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshbohde/example"
)

type Server struct {
	AddressService example.AddressService
	UserService    example.UserService
}

func (s *Server) Run(port int) {
	engine := gin.Default()

	users := UserHandler{UserService: s.UserService}

	engine.GET("/users/:id", users.Get)

	engine.Run(":" + strconv.Itoa(port))

}
