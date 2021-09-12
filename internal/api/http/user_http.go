package http

import (
	"base_service/internal/application/dtos"
	"base_service/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) userHandler(c *gin.Context) {
	var user dtos.UserDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userEntities := &entities.User{
		Username: user.Username,
	}
	userResponse, err := s.handler.GetUser(userEntities)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userResponse)
}
