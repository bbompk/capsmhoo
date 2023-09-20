package user

import (
	"github.com/gin-gonic/gin"
)

// Define Http Endpoints Here
func ProvideRouter(r *gin.Engine, h *Handler) {
	// GET a user by ID
	r.GET("/user/:id", h.GetUserByID)

	// GET all users
	r.GET("/user", h.GetUser)

	// POST (create) a new user
	r.POST("/user", h.CreateUser)

	// PUT (update) a user by ID
	r.PUT("/user/:id", h.UpdateUserByID)

	// DELETE a user by ID
	r.DELETE("/user/:id", h.DeleteUserByID)
}
