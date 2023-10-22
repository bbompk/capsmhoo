package http_handler

import (
	restClient "capsmhoo/mono/api-gateway/client_rest"

	"github.com/gin-gonic/gin"
	// "capsmhoo/mono/api-gateway/model"
)

type StudentHandler struct {
	studentClientRest restClient.StudentClientRest
}

type IStudentHandler interface {
	GetStudentByID(c *gin.Context)
	GetAllStudents(c *gin.Context)
	CreateStudent(c *gin.Context)
	UpdateStudentByID(c *gin.Context)
	DeleteStudentByID(c *gin.Context)
	DeleteStudentAll(c *gin.Context)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
}

func (h *StudentHandler) GetAllStudents(c *gin.Context) {
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
}

func (h *StudentHandler) UpdateStudentByID(c *gin.Context) {
}

func (h *StudentHandler) DeleteStudentByID(c *gin.Context) {
}

func (h *StudentHandler) DeleteStudentAll(c *gin.Context) {
}

func ProvideStudentHandler(studentClientRest restClient.StudentClientRest) *StudentHandler {
	return &StudentHandler{
		studentClientRest: studentClientRest,
	}
}
