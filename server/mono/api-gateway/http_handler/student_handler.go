package http_handler

import (
	"github.com/gin-gonic/gin"
	restClient "capsmhoo/mono/api-gateway/client_rest"
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

func (h *TeamHandler) GetStudentByID(c *gin.Context) {
}

func (h *TeamHandler) GetAllStudents(c *gin.Context) {
}

func (h *TeamHandler) CreateStudent(c *gin.Context) {
}

func (h *TeamHandler) UpdateStudentByID(c *gin.Context) {
}

func (h *TeamHandler) DeleteStudentByID(c *gin.Context) {
}

func (h *TeamHandler) DeleteStudentAll(c *gin.Context) {
}

func ProvideStudentHandler(studentClientRest restClient.StudentClientRest) *StudentHandler {
	return &StudentHandler{
		studentClientRest: studentClientRest,
	}
}