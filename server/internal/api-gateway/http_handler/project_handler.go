package http_handler

import (
	grpcClient "capsmhoo/internal/api-gateway/client_grpc"
	"capsmhoo/internal/api-gateway/model"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectClientgRPC grpcClient.ProjectgRPCClient
}

type IProjectHandler interface {
	GetAllProjects(c *gin.Context)
	GetProjectByID(c *gin.Context)
	GetProjectByTeamID(c *gin.Context)
	GetProjectByProfessorID(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProjectByID(c *gin.Context)
	DeleteProjectByID(c *gin.Context)
	// AddStudentToProject(c *gin.Context)
	// RemoveStudentFromProject(c *gin.Context)
	GetProjectRequestByProjectID(c *gin.Context)
	CreateProjectRequest(c *gin.Context)
	AcceptProjectRequest(c *gin.Context)
	RejectProjectRequest(c *gin.Context)
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.projectClientgRPC.GetAllProjects(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": projects,
	})
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	project, err := h.projectClientgRPC.GetProjectByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": project,
	})
}

func (h *ProjectHandler) GetProjectByTeamID(c *gin.Context) {
	id := c.Param("id")
	project, err := h.projectClientgRPC.GetProjectByTeamID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "200",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": project,
	})
}

func (h *ProjectHandler) GetProjectByProfessorID(c *gin.Context) {
	id := c.Param("id")
	projects, err := h.projectClientgRPC.GetProjectByProfessorID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": projects,
	})
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	createdProject, err := h.projectClientgRPC.CreateProject(c, &project)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": createdProject,
	})
}

func (h *ProjectHandler) UpdateProjectByID(c *gin.Context) {
	id := c.Param("id")
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	updatedProject, err := h.projectClientgRPC.UpdateProjectByID(c, id, &project)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": updatedProject,
	})
}

func (h *ProjectHandler) DeleteProjectByID(c *gin.Context) {
	id := c.Param("id")
	deletedProject, err := h.projectClientgRPC.DeleteProjectByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": deletedProject,
	})
}

// func (h *TeamHandler) AddStudentToTeam(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"code":    "200",
// 		"message": "AddStudentToTeam",
// 	})
// }

// func (h *TeamHandler) RemoveStudentFromTeam(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"code":    "200",
// 		"message": "RemoveStudentFromTeam",
// 	})
// }

func (h *ProjectHandler) GetProjectRequestByProjectID(c *gin.Context) {
	id := c.Param("id")
	projectRequests, err := h.projectClientgRPC.GetProjectRequestByProjectID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": projectRequests,
	})
}

func (h *ProjectHandler) CreateProjectRequest(c *gin.Context) {
	var projectRequest model.ProjectRequest
	if err := c.ShouldBindJSON(&projectRequest); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	createdProjectRequest, err := h.projectClientgRPC.CreateProjectRequest(c, &projectRequest)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": createdProjectRequest,
	})
}

func (h *ProjectHandler) AcceptProjectRequest(c *gin.Context) {
	id := c.Param("id")
	_, err := h.projectClientgRPC.AcceptProjectRequest(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Project Request Accepted",
	})
}

func (h *ProjectHandler) RejectProjectRequest(c *gin.Context) {
	id := c.Param("id")
	_, err := h.projectClientgRPC.RejectProjectRequest(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Project Request Rejected",
	})
}

func ProvideProjectHandler(projectClientgRPC grpcClient.ProjectgRPCClient) *ProjectHandler {
	return &ProjectHandler{
		projectClientgRPC: projectClientgRPC,
	}
}
