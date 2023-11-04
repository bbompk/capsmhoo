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
	r.POST("/auth/login", userHandler.Login)

	// user service
	r.GET("/user/:id", userHandler.GetUserByID)
	r.GET("/user", userHandler.GetAllUsers)
	r.POST("/user", userHandler.CreateUser)
	r.PUT("/user/:id", userHandler.UpdateUserByID)
	r.DELETE("/user/:id", userHandler.DeleteUserByID)

	// student service
	r.GET("/student/userId/:id", studentHandler.GetStudentByUserID)
	r.GET("/student/teamId/:id", studentHandler.GetAllStudentsByTeamID)
	r.GET("/student/:id", studentHandler.GetStudentByID)
	r.GET("/student", studentHandler.GetAllStudents)
	r.POST("/student", studentHandler.CreateStudent)
	r.PUT("/student/:id", studentHandler.UpdateStudentByID)
	r.DELETE("/student/:id", studentHandler.DeleteStudentByID)
	r.DELETE("/student", studentHandler.DeleteStudentAll)

	// professor service
	r.GET("/professor/userId/:id", professorHandler.GetProfessorByUserID)
	r.GET("/professor/:id", professorHandler.GetProfessorByID)
	r.GET("/professor", professorHandler.GetAllProfessors)
	r.POST("/professor", professorHandler.CreateProfessor)
	r.PUT("/professor/:id", professorHandler.UpdateProfessorByID)
	r.DELETE("/professor/:id", professorHandler.DeleteProfessorByID)
	r.DELETE("/professor", professorHandler.DeleteProfessorAll)

	// noti service
	notiRoute := r.Group("/noti")
	notiRoute.GET("/:id", notiHandler.GetAllNotiByUserId)
	notiRoute.POST("/:id", notiHandler.ReadNoti)

	// team service
	r.GET("/team", teamHandler.GetAllTeams)
	r.GET("/team/:id", teamHandler.GetTeamByID)
	r.POST("/team", teamHandler.CreateTeam)
	r.PUT("/team/:id", teamHandler.UpdateTeamByID)
	r.DELETE("/team/:id", teamHandler.DeleteTeamByID)
	r.POST("/team/add-student/:id", teamHandler.AddStudentToTeam)
	r.POST("/team/remove-student/:id", teamHandler.RemoveStudentFromTeam)

	teamJoinRequestRoute := r.Group("/team-join-request")
	teamJoinRequestRoute.GET("", teamJoinRequestHandler.GetAllJoinRequests)
	teamJoinRequestRoute.GET("/:id", teamJoinRequestHandler.GetJoinRequestByID)
	teamJoinRequestRoute.GET("/teamid/:id", teamJoinRequestHandler.GetJoinRequestByTeamID)
	teamJoinRequestRoute.POST("", teamJoinRequestHandler.CreateJoinRequest)
	teamJoinRequestRoute.PUT("/:id", teamJoinRequestHandler.UpdateJoinRequest)
	teamJoinRequestRoute.DELETE("/:id", teamJoinRequestHandler.DeleteJoinRequest)
	teamJoinRequestRoute.POST("/approve/:id", teamJoinRequestHandler.ApproveJoinRequest)
	teamJoinRequestRoute.POST("/decline/:id", teamJoinRequestHandler.DeclineJoinRequest)

	// project service
	r.GET("/project", projectHandler.GetAllProjects)
	r.GET("/project/:id", projectHandler.GetProjectByID)
	r.POST("/project", projectHandler.CreateProject)
	r.PUT("/project/:id", projectHandler.UpdateProjectByID)
	r.DELETE("/project/:id", projectHandler.DeleteProjectByID)
	r.GET("/project-request/projectid/:id", projectHandler.GetProjectRequestByProjectID)
	r.POST("/project-request", projectHandler.CreateProjectRequest)
	r.POST("/project-request/accept/:id", projectHandler.AcceptProjectRequest)
	r.POST("/project-request/reject/:id", projectHandler.RejectProjectRequest)
}
