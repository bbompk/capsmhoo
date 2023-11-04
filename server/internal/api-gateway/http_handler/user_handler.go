package http_handler

import (
	"net/http"

	restClient "capsmhoo/internal/api-gateway/client_rest"
	"capsmhoo/internal/api-gateway/model"

	"github.com/gin-gonic/gin"
	// "capsmhoo/internal/api-gateway/model"
)

type UserHandler struct {
	userClientRest restClient.UserClientRest
}

type IUserHandler interface {
	GetUserByID(c *gin.Context)
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUserByID(c *gin.Context)
	DeleteUserByID(c *gin.Context)
	DeleteUserAll(c *gin.Context)
	Login(c *gin.Context)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userClientRest.GetUserByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userClientRest.GetAllUsers() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var params model.UserRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	user, err := h.userClientRest.CreateUser(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}

func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var params model.UserRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	user, err := h.userClientRest.UpdateUserByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userClientRest.DeleteUserByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}

func (h *UserHandler) DeleteUserAll(c *gin.Context) {
}

func (h *UserHandler) Login(c *gin.Context) {
	var params model.UserRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	resp, err := h.userClientRest.Login(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	resp.User.Password = ""
	c.JSON(200, gin.H{
		"code": "200",
		"data": resp,
	})
}

func ProvideUserHandler(userClientRest restClient.UserClientRest) *UserHandler {
	return &UserHandler{
		userClientRest: userClientRest,
	}
}
