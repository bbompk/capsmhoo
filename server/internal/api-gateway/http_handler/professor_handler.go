package http_handler

import (
	restClient "capsmhoo/internal/api-gateway/client_rest"

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
}

func (h *ProfessorHandler) GetProfessorByID(c *gin.Context) {
}

func (h *ProfessorHandler) GetAllProfessors(c *gin.Context) {
}

func (h *ProfessorHandler) CreateProfessor(c *gin.Context) {
}

func (h *ProfessorHandler) UpdateProfessorByID(c *gin.Context) {
}

func (h *ProfessorHandler) DeleteProfessorByID(c *gin.Context) {
}

func (h *ProfessorHandler) DeleteProfessorAll(c *gin.Context) {
}

func ProvideProfessorHandler(professorClientRest restClient.ProfessorClientRest) *ProfessorHandler {
	return &ProfessorHandler{
		professorClientRest: professorClientRest,
	}
}
