package http_handler

import (
	"github.com/gin-gonic/gin"
	restClient "capsmhoo/mono/api-gateway/client_rest"
	// "capsmhoo/mono/api-gateway/model"
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
}

func (h *TeamHandler) GetProfessorByID(c *gin.Context) {
}

func (h *TeamHandler) GetAllProfessors(c *gin.Context) {
}

func (h *TeamHandler) CreateProfessor(c *gin.Context) {
}

func (h *TeamHandler) UpdateProfessorByID(c *gin.Context) {
}

func (h *TeamHandler) DeleteProfessorByID(c *gin.Context) {
}

func (h *TeamHandler) DeleteProfessorAll(c *gin.Context) {
}

func ProvideProfessorHandler(professorClientRest restClient.ProfessorClientRest) *ProfessorHandler {
	return &ProfessorHandler{
		professorClientRest: professorClientRest,
	}
}