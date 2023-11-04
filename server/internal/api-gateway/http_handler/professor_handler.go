package http_handler

import (
	"net/http"

	restClient "capsmhoo/internal/api-gateway/client_rest"
	"capsmhoo/internal/api-gateway/model"

	"github.com/gin-gonic/gin"
	// "capsmhoo/internal/api-gateway/model"
)

type ProfessorHandler struct {
	professorClientRest restClient.ProfessorClientRest
}

type IProfessorHandler interface {
	GetProfessorByID(c *gin.Context)
	GetAllProfessors(c *gin.Context)
	CreateProfessor(c *gin.Context)
	UpdateProfessorByID(c *gin.Context)
	DeleteProfessorByID(c *gin.Context)
	DeleteProfessorAll(c *gin.Context)
	GetProfessorByUserID(c *gin.Context)
}

func (h *ProfessorHandler) GetProfessorByID(c *gin.Context) {
	id := c.Param("id")
	professor, err := h.professorClientRest.GetProfessorByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": professor})
}

func (h *ProfessorHandler) GetAllProfessors(c *gin.Context) {
	professors, err := h.professorClientRest.GetAllProfessors() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": professors})
}

func (h *ProfessorHandler) CreateProfessor(c *gin.Context) {
	var params model.ProfessorRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	professor, err := h.professorClientRest.CreateProfessor(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": professor,
	})
}

func (h *ProfessorHandler) UpdateProfessorByID(c *gin.Context) {
	id := c.Param("id")
	var params model.ProfessorRequestBody
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	professor, err := h.professorClientRest.UpdateProfessorByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": professor,
	})
}

func (h *ProfessorHandler) DeleteProfessorByID(c *gin.Context) {
	id := c.Param("id")
	professor, err := h.professorClientRest.DeleteProfessorByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": professor,
	})
}

func (h *ProfessorHandler) GetProfessorByUserID(c *gin.Context) {
	id := c.Param("id")
	professor, err := h.professorClientRest.GetProfessorByUserID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": professor,
	})
}

func (h *ProfessorHandler) DeleteProfessorAll(c *gin.Context) {
}

func ProvideProfessorHandler(professorClientRest restClient.ProfessorClientRest) *ProfessorHandler {
	return &ProfessorHandler{
		professorClientRest: professorClientRest,
	}
}
