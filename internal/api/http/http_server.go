package http

import (
	app "base_service/internal/application"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handler *app.UserHandler
}

func NewServer(handler *app.UserHandler) *Server {
	return &Server{handler: handler}
}

func Listen(port string, s *Server) {
	r := gin.Default()
	r.POST("/users", s.userHandler)
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
