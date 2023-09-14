package user

import (
	"capsmhoo/common"

	"github.com/gin-gonic/gin"
)

// Define Dependencies
type Handler struct {
	repo UserRepository
}

// Define what this will do
type UserHttpHandler interface {
	GetUser(c *gin.Context)
}

func (h *Handler) GetUser(c *gin.Context) {
	user := h.repo.GetUser()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: user,
	})
}

// Dependency Injection
func ProvideHandler(repo UserRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}
