package team

import (
	"github.com/gin-gonic/gin"
)

// Define Http Endpoints Here
func ProvideRouter(r *gin.Engine, h *TeamHandler) {
	// GET all teams
	r.GET("/team", h.GetTeams)

	// GET a team by ID
	r.GET("/team/:id", h.GetTeamByID)

	// POST (create) a new team
	r.POST("/team", h.CreateTeam)

	// PUT (update) a team by ID
	r.PUT("/team/:id", h.UpdateTeamByID)

	// DELETE a team by ID
	r.DELETE("/team/:id", h.DeleteTeamByID)

	//DELETE all teams
	r.DELETE("/team", h.DeleteAll)
}
