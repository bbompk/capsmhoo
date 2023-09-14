package user

import (
	"github.com/gin-gonic/gin"
)

// Define Http Endpoints Here
func ProvideRouter(r *gin.Engine, h *Handler) {
	r.GET("/user", h.GetUser)
}
