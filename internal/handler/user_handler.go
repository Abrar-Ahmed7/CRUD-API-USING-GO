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
	router.POST("/user", uh.Create)
}

func (uh userHandler) Create(c *gin.Context) {

	type UserReqBody struct {
		Name    string `json:"name" binding:"required,min=3,max=10"`
		EmailId string `json:"email_id" binding:"required,min=5,max=100"`
	}

	var req UserReqBody
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name or email can't be empty",
		})
		return
	}

	name := req.Name
	emailId := req.EmailId
	res := uh.userService.Save(name, emailId)
	c.JSON(http.StatusOK, gin.H{
		"message":  res,
		"name":     name,
		"email_id": emailId,
	})
}
