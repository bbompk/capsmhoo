package user

import (
	"capsmhoo/common"

	"github.com/gin-gonic/gin"
)

// Define Dependencies
type UserHandler struct {
	repo UserRepository
}

// Define what this will do
type UserHttpHandler interface {
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUserByID(c *gin.Context)
	DeleteUserByID(c *gin.Context)
	DeleteAll(c *gin.Context)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user := h.repo.GetUsers()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: user,
	})
}
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.repo.GetUserByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: user,
	})
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	// user := h.repo.CreateUser()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: createdUser,
	})
}
func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	updatedUser, err := h.repo.UpdateUserByID(id, user)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: updatedUser,
	})
}
func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.DeleteUserByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: "",
	})
}
func (h *UserHandler) DeleteAll(c *gin.Context) {
	h.repo.DeleteAll()
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: "",
	})
}

// Dependency Injection
func ProvideUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}
