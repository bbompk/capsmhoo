package http_handler

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	teamHandler ITeamHandler,
) {
	r.GET("/team", teamHandler.GetAllTeams)
	r.GET("/team/:id", teamHandler.GetTeamByID)
	r.POST("/team", teamHandler.CreateTeam)
	r.PUT("/team/:id", teamHandler.UpdateTeamByID)
	r.DELETE("/team/:id", teamHandler.DeleteTeamByID)
	r.POST("/team/add-student/:id", teamHandler.AddStudentToTeam)
	r.POST("/team/remove-student/:id", teamHandler.RemoveStudentFromTeam)
}
