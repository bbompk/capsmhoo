package user

import (
	"capsmhoo/common"

	"github.com/gin-gonic/gin"
)

// Define Dependencies
type ProfessorHandler struct {
	repo     ProfessorRepository
	userrepo UserRepository
}

// Define what this will do
type ProfessorHttpHandler interface {
	GetProfessors(c *gin.Context)
	GetProfessorByID(c *gin.Context)
	CreateProfessor(c *gin.Context)
	UpdateProfessorByID(c *gin.Context)
	DeleteProfessorByID(c *gin.Context)
}

func (h *ProfessorHandler) GetProfessor(c *gin.Context) {
	professor := h.repo.GetProfessors()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: professor,
	})
}
func (h *ProfessorHandler) GetProfessorByID(c *gin.Context) {
	id := c.Param("id")
	professor, err := h.repo.GetProfessorByID(id)
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
		Data: professor,
	})
}
func (h *ProfessorHandler) CreateProfessor(c *gin.Context) {
	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Profile  string `json:"profile"`
	}
	var params Params
	var professor Professor
	var user User
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: "Couldn't bind input to json",
		})
		return
	}
	user.Email = params.Email
	user.Password = params.Password
	professor.Name = params.Name
	professor.Profile = params.Profile
	createdUser, err := h.userrepo.CreateUser(user)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: "User cannot be created",
		})
		return
	}
	// time.Sleep(time.Second)
	professor.UserID = createdUser.ID
	createdProfessor, err := h.repo.CreateProfessor(professor)
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
		Data: createdProfessor,
	})
}
func (h *ProfessorHandler) UpdateProfessorByID(c *gin.Context) {
	id := c.Param("id")
	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Profile  string `json:"profile"`
	}
	var params Params
	var professor Professor
	var user User
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: "Couldn't bind input to json",
		})
		return
	}
	if params.Email != "" {
		user.Email = params.Email
	}
	if params.Password != "" {
		user.Password = params.Password
	}
	if params.Name != "" {
		professor.Name = params.Name
	}
	if params.Profile != "" {
		professor.Profile = params.Profile
	}
	professor.ID = id
	professorr, err := h.repo.GetProfessorByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	user.ID = professorr.UserID
	_, errr := h.userrepo.UpdateUserByID(user.ID, user)
	if errr != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	updatedProfessor, err := h.repo.UpdateProfessorByID(id, professor)
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
		Data: updatedProfessor,
	})
}
func (h *ProfessorHandler) DeleteProfessorByID(c *gin.Context) {
	id := c.Param("id")
	professor, err := h.repo.GetProfessorByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	errrr := h.repo.DeleteProfessorByID(id)
	if errrr != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	errr := h.userrepo.DeleteUserByID(professor.UserID)
	if errr != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: professor,
	})
}

// Dependency Injection
func ProvideProfessorHandler(repo ProfessorRepository, userrepo UserRepository) *ProfessorHandler {
	return &ProfessorHandler{
		repo:     repo,
		userrepo: userrepo,
	}
}
