package http_handler

import (
	"capsmhoo/mono/api-gateway/model"
	"github.com/gin-gonic/gin"
	"capsmhoo/mono/api-gateway/client_rest" 
	"net/http"
)

type IUserHandler interface {
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUserByID(c *gin.Context)
	DeleteUserByID(c *gin.Context)
	// DeleteAll(c *gin.Context)
}

// NewUserHandler initializes a new handler with the given REST client.
func NewUserHandler(client client_rest.UserRestClient) *UserHandler {
	return &UserHandler{
		userRestClient: client,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userRestClient.GetAllUsers() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userRestClient.GetUser(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User // Assuming the user model is consistent across services
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.userRestClient.CreateUser(&user) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response) // Adjust based on how you want the response
}

func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.userRestClient.UpdateUser(id, &user) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response) // Adjust based on how you want the response
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	response, err := h.userRestClient.DeleteUser(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response) // Adjust based on how you want the response
}


