package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"request-handler-unit-test-example/users/dto"
)

type RequestHandler struct {
	controller Controller
}

func (h RequestHandler) GetUsers(c *gin.Context) {
	res, err := h.controller.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.MessageResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) CreateUser(c *gin.Context) {
	req := dto.CreateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{Message: err.Error()})
		return
	}

	res, err := h.controller.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.MessageResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
