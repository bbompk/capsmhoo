package http_handler

import (
	grpcClient "capsmhoo/internal/api-gateway/client_grpc"
	"capsmhoo/internal/api-gateway/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type TeamHandler struct {
	teamClientgRPC grpcClient.TeamgRPCClient
}

type TeamJoinRequestHandler struct {
	teamJoinRequestClientgRPC grpcClient.TeamJoinRequestgRPCClient
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

type ITeamJoinRequestHandler interface {
	GetAllJoinRequests(c *gin.Context)
	GetJoinRequestByID(c *gin.Context)
	GetJoinRequestByTeamID(c *gin.Context)
	CreateJoinRequest(c *gin.Context)
	UpdateJoinRequest(c *gin.Context)
	DeleteJoinRequest(c *gin.Context)
	ApproveJoinRequest(c *gin.Context)
	DeclineJoinRequest(c *gin.Context)
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
		"code": "200",
		"data": teams,
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
		"data": team,
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
		"data": createdTeam,
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
		"data": updatedTeam,
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
		"data": deletedTeam,
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

func (h *TeamJoinRequestHandler) GetAllJoinRequests(c *gin.Context) {
	requests, err := h.teamJoinRequestClientgRPC.GetAllJoinRequests(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": requests,
	})
}

func (h *TeamJoinRequestHandler) GetJoinRequestByID(c *gin.Context) {
	id := c.Param("id")
	request, err := h.teamJoinRequestClientgRPC.GetJoinRequestByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": request,
	})
}

func (h *TeamJoinRequestHandler) GetJoinRequestByTeamID(c *gin.Context) {
	id := c.Param("id")
	requests, err := h.teamJoinRequestClientgRPC.GetJoinRequestByTeamID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": requests,
	})
}

func (h *TeamJoinRequestHandler) CreateJoinRequest(c *gin.Context) {
	var request model.TeamJoinRequest

	user, _ := c.Get("user")
	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	request.StudentID = userID

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	createdRequest, err := h.teamJoinRequestClientgRPC.CreateJoinRequest(c, &request)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": createdRequest,
	})
}

func (h *TeamJoinRequestHandler) UpdateJoinRequest(c *gin.Context) {
	id := c.Param("id")
	var request model.TeamJoinRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	updatedRequest, err := h.teamJoinRequestClientgRPC.UpdateJoinRequest(c, id, &request)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": updatedRequest,
	})
}

func (h *TeamJoinRequestHandler) DeleteJoinRequest(c *gin.Context) {
	id := c.Param("id")
	_, err := h.teamJoinRequestClientgRPC.DeleteJoinRequest(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	// If needed, use 'deletedRequest' here. Otherwise, you can ignore it.
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Successfully deleted",
	})
}

func (h *TeamJoinRequestHandler) ApproveJoinRequest(c *gin.Context) {
	id := c.Param("id")
	_, err := h.teamJoinRequestClientgRPC.ApproveJoinRequest(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Successfully approved",
	})
}

func (h *TeamJoinRequestHandler) DeclineJoinRequest(c *gin.Context) {
	id := c.Param("id")
	_, err := h.teamJoinRequestClientgRPC.DeclineJoinRequest(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Successfully declined",
	})
}

func ProvideTeamHandler(teamClientgRPC grpcClient.TeamgRPCClient) *TeamHandler {
	return &TeamHandler{
		teamClientgRPC: teamClientgRPC,
	}
}

func ProvideTeamJoinRequestHandler(TeamJoinRequestClientgRPC grpcClient.TeamJoinRequestgRPCClient) *TeamJoinRequestHandler {
	return &TeamJoinRequestHandler{
		teamJoinRequestClientgRPC: TeamJoinRequestClientgRPC,
	}
}
