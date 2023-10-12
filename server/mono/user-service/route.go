package user

import (
	"github.com/gin-gonic/gin"
)

// Define Http Endpoints Here
func ProvideRouter(r *gin.Engine, h *UserHandler, hh *StudentHandler, hhh *ProfessorHandler) {
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
	r.DELETE("/user", h.DeleteAll)

	// GET a student by ID
	r.GET("/student/:id", hh.GetStudentByID)

	// GET all students
	r.GET("/student", hh.GetStudent)

	// POST (create) a new student
	r.POST("/student", hh.CreateStudent)

	// PUT (update) a student by ID
	r.PUT("/student/:id", hh.UpdateStudentByID)

	// DELETE a student by ID
	r.DELETE("/student/:id", hh.DeleteStudentByID)

	// GET a professor by ID
	r.GET("/professor/:id", hhh.GetProfessorByID)

	// GET all professors
	r.GET("/professor", hhh.GetProfessor)

	// POST (create) a new professor
	r.POST("/professor", hhh.CreateProfessor)

	// PUT (update) a professor by ID
	r.PUT("/professor/:id", hhh.UpdateProfessorByID)

	// DELETE a professor by ID
	r.DELETE("/professor/:id", hhh.DeleteProfessorByID)
}
