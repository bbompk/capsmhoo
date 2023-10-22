package http_handler

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	teamHandler ITeamHandler,
	teamJoinRequestHandler ITeamJoinRequestHandler,
	notiHandler INotiHandler,
) {
	r.GET("/team", teamHandler.GetAllTeams)
	r.GET("/team/:id", teamHandler.GetTeamByID)
	r.POST("/team", teamHandler.CreateTeam)
	r.PUT("/team/:id", teamHandler.UpdateTeamByID)
	r.DELETE("/team/:id", teamHandler.DeleteTeamByID)
	r.POST("/team/add-student/:id", teamHandler.AddStudentToTeam)
	r.POST("/team/remove-student/:id", teamHandler.RemoveStudentFromTeam)

	notiRoute := r.Group("/noti")
	notiRoute.GET("/:id", notiHandler.GetAllNotiByUserId)
	notiRoute.POST("/:id", notiHandler.ReadNoti)
	teamJoinRequestRoute := r.Group("/team-join-request")
	teamJoinRequestRoute.GET("", teamJoinRequestHandler.GetAllJoinRequests)
	teamJoinRequestRoute.GET("/:id", teamJoinRequestHandler.GetJoinRequestByID)
	teamJoinRequestRoute.POST("", teamJoinRequestHandler.CreateJoinRequest)
	teamJoinRequestRoute.PUT("/:id", teamJoinRequestHandler.UpdateJoinRequest)
	teamJoinRequestRoute.DELETE("/:id", teamJoinRequestHandler.DeleteJoinRequest)
	teamJoinRequestRoute.POST("/approve/:id", teamJoinRequestHandler.ApproveJoinRequest)
	teamJoinRequestRoute.POST("/decline/:id", teamJoinRequestHandler.DeclineJoinRequest)
}
