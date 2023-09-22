package team

import (
	"capsmhoo/common"

	"github.com/gin-gonic/gin"
)

// Define Dependencies
type TeamHandler struct {
	repo TeamRepository
}

// Define what this will do
type TeamHttpHandler interface {
	GetTeams(c *gin.Context)
	GetTeamByID(c *gin.Context)
	CreateTeam(c *gin.Context)
	UpdateTeamByID(c *gin.Context)
	DeleteTeamByID(c *gin.Context)
	DeleteAll(c *gin.Context)
}

func (h *TeamHandler) GetTeams(c *gin.Context) {
	teams := h.repo.GetTeams()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: teams,
	})
}
func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	id := c.Param("id")
	team, err := h.repo.GetTeamByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code:  "400",
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: team,
	})
}
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: "Couldn't bind input to json",
		})
		return
	}
	createdTeam, err := h.repo.CreateTeam(team)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: "Team cannot be created",
		})
		return
	}

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: createdTeam,
	})
}

func (h *TeamHandler) UpdateTeamByID(c *gin.Context) {
	id := c.Param("id")
	var team Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	updatedTeam, err := h.repo.UpdateTeamByID(id, team)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: updatedTeam,
	})
}
func (h *TeamHandler) DeleteTeamByID(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.DeleteTeamByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: "",
	})
}
func (h *TeamHandler) DeleteAll(c *gin.Context) {
	h.repo.DeleteAll()
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: "",
	})
}

// Dependency Injection
func ProvideTeamHandler(repo TeamRepository) *TeamHandler {
	return &TeamHandler{
		repo: repo,
	}
}
