package handler

import (
	"net/http"

	"github.com/Abrar-Ahmed7/rest-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(c *gin.Context)
	RegisterRoutes(router *gin.RouterGroup)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh userHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/user", uh.Create)
}

func (uh userHandler) Create(c *gin.Context) {
	res := uh.userService.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
