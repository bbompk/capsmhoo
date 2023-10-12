package http_handler

import (
	grpcClient "capsmhoo/mono/api-gateway/client_grpc"
	"capsmhoo/mono/api-gateway/model"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamClientgRPC grpcClient.TeamgRPCClient
}

type ITeamHandler interface {
	GetAllTeams(c *gin.Context)
	GetTeamByID(c *gin.Context)
	CreateTeam(c *gin.Context)
	UpdateTeamByID(c *gin.Context)
	DeleteTeamByID(c *gin.Context)
	AddStudentToTeam(c *gin.Context)
	RemoveStudentFromTeam(c *gin.Context)
}

func (h *TeamHandler) GetAllTeams(c *gin.Context) {
	teams, err := h.teamClientgRPC.GetAllTeams(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"teams": teams,
	})
}

func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	id := c.Param("id")
	team, err := h.teamClientgRPC.GetTeamByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"team": team,
	})
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team model.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	createdTeam, err := h.teamClientgRPC.CreateTeam(c, &team)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"team": createdTeam,
	})
}

func (h *TeamHandler) UpdateTeamByID(c *gin.Context) {
	id := c.Param("id")
	var team model.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	updatedTeam, err := h.teamClientgRPC.UpdateTeamByID(c, id, &team)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"team": updatedTeam,
	})
}

func (h *TeamHandler) DeleteTeamByID(c *gin.Context) {
	id := c.Param("id")
	deletedTeam, err := h.teamClientgRPC.DeleteTeamByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"team": deletedTeam,
	})
}

func (h *TeamHandler) AddStudentToTeam(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "AddStudentToTeam",
	})
}

func (h *TeamHandler) RemoveStudentFromTeam(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "RemoveStudentFromTeam",
	})
}

func ProvideTeamHandler(teamClientgRPC grpcClient.TeamgRPCClient) *TeamHandler {
	return &TeamHandler{
		teamClientgRPC: teamClientgRPC,
	}
}
