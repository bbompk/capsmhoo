package http_handler

import (
	"net/http"

	restClient "capsmhoo/internal/api-gateway/client_rest"
	"capsmhoo/internal/api-gateway/model"

	"github.com/gin-gonic/gin"
	// "capsmhoo/internal/api-gateway/model"
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
	GetAllStudentsByTeamID(c *gin.Context)
	GetStudentByUserID(c *gin.Context)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.studentClientRest.GetStudentByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "500",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": student})
}

func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := h.studentClientRest.GetAllStudents() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "500",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": students})
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var params model.StudentRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	student, err := h.studentClientRest.CreateStudent(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": student,
	})
}

func (h *StudentHandler) UpdateStudentByID(c *gin.Context) {
	id := c.Param("id")
	var params model.StudentRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	student, err := h.studentClientRest.UpdateStudentByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": student,
	})
}

func (h *StudentHandler) DeleteStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.studentClientRest.DeleteStudentByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": student,
	})
}

func (h *StudentHandler) GetAllStudentsByTeamID(c *gin.Context) {
	id := c.Param("id")
	students, err := h.studentClientRest.GetAllStudentsByTeamID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "500",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": students,
	})
}

func (h *StudentHandler) GetStudentByUserID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.studentClientRest.GetStudentByUserID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "500",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": student,
	})
}

func (h *StudentHandler) DeleteStudentAll(c *gin.Context) {
}

func ProvideStudentHandler(studentClientRest restClient.StudentClientRest) *StudentHandler {
	return &StudentHandler{
		studentClientRest: studentClientRest,
	}
}
