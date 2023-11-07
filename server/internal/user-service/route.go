package user

import (
	"github.com/gin-gonic/gin"
)

// Define Http Endpoints Here
func ProvideRouter(r *gin.Engine, h *UserHandler, hh *StudentHandler, hhh *ProfessorHandler) {

	// login service
	// r.POST("/register", h.SignUpUser)
	r.POST("/login", h.SignInUser)
	// r.POST("/logout", middleware.JwtAuthentication, h.SignOutUser)

	r.GET("/user/:id", h.GetUserByID)
	r.GET("/user", h.GetUser)
	r.POST("/user", h.CreateUser)
	// r.PUT("/user/:id", h.UpdateUserByID)
	r.PUT("/user/:id", h.UpdateUserByID)
	r.DELETE("/user/:id", h.DeleteUserByID)
	r.DELETE("/user", h.DeleteAll)

	r.GET("/student/userId/:user_id", hh.GetStudentByUserID)
	r.GET("/student/teamId/:team_id", hh.GetStudentByTeamID)
	r.GET("/student/:id", hh.GetStudentByID)
	r.GET("/student", hh.GetStudent)
	r.POST("/student", hh.CreateStudent)
	r.PUT("/student/:id", hh.UpdateStudentByID)
	r.DELETE("/student/:id", hh.DeleteStudentByID)

	r.GET("/professor/userId/:user_id", hhh.GetProfessorByUserID)
	r.GET("/professor/:id", hhh.GetProfessorByID)
	r.GET("/professor", hhh.GetProfessor)
	r.POST("/professor", hhh.CreateProfessor)
	r.PUT("/professor/:id", hhh.UpdateProfessorByID)
	r.DELETE("/professor/:id", hhh.DeleteProfessorByID)
}
