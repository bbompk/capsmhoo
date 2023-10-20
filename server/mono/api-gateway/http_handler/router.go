package http_handler

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	teamHandler ITeamHandler,
	teamJoinRequestHandler ITeamJoinRequestHandler,
	userHandler IUserHandler,
	studentHandler IStudentHandler,
	professorHandler IProfessorHandler,
	projectHandler IProjectHandler,
	notiHandler INotiHandler,
) {
	// team service
	r.GET("/team", teamHandler.GetAllTeams)
	r.GET("/team/:id", teamHandler.GetTeamByID)
	r.POST("/team", teamHandler.CreateTeam)
	r.PUT("/team/:id", teamHandler.UpdateTeamByID)
	r.DELETE("/team/:id", teamHandler.DeleteTeamByID)
	r.POST("/team/add-student/:id", teamHandler.AddStudentToTeam)
	r.POST("/team/remove-student/:id", teamHandler.RemoveStudentFromTeam)

	// user service
	r.GET("/user/:id", userHandler.GetUserByID)
	r.GET("/user", userHandler.GetUser)
	r.POST("/user", userHandler.CreateUser)
	r.PUT("/user/:id", userHandler.UpdateUserByID)
	r.DELETE("/user/:id", userHandler.DeleteUserByID)
	r.DELETE("/user", userHandler.DeleteAll)

	r.GET("/student/:id", studentHandler.GetStudentByID)
	r.GET("/student", studentHandler.GetStudent)
	r.POST("/student", studentHandler.CreateStudent)
	r.PUT("/student/:id", studentHandler.UpdateStudentByID)
	r.DELETE("/student/:id", studentHandler.DeleteStudentByID)
	r.DELETE("/student", studentHandler.DeleteStudentAll)

	r.GET("/professor/:id", professorHandler.GetProfessorByID)
	r.GET("/professor", professorHandler.GetProfessor)
	r.POST("/professor", professorHandler.CreateProfessor)
	r.PUT("/professor/:id", professorHandler.UpdateProfessorByID)
	r.DELETE("/professor/:id", professorHandler.DeleteProfessorByID)
	r.DELETE("/professor", professorHandler.DeleteProfessorAll)

	// project service
	// todo: wait for project service

	// noti service
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
