package user

import (
	"capsmhoo/common"

	"github.com/gin-gonic/gin"
)

// Define Dependencies
type StudentHandler struct {
	repo     StudentRepository
	userrepo UserRepository
}

// Define what this will do
type StudentHttpHandler interface {
	GetStudents(c *gin.Context)
	GetStudentByID(c *gin.Context)
	CreateStudent(c *gin.Context)
	UpdateStudentByID(c *gin.Context)
	DeleteStudentByID(c *gin.Context)
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	student := h.repo.GetStudents()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: student,
	})
}
func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.repo.GetStudentByID(id)
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
		Data: student,
	})
}
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	var params Params
	var student Student
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
	student.Name = params.Name
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
	student.UserID = createdUser.ID
	createdStudent, err := h.repo.CreateStudent(student)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	// student := h.repo.CreateStudent()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: createdStudent,
	})
}
func (h *StudentHandler) UpdateStudentByID(c *gin.Context) {
	id := c.Param("id")

	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		// TeamID   string `json:"team_id"`
	}
	var params Params
	var student Student
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
		student.Name = params.Name
	}
	// if params.TeamID != "" {
	//     student.TeamID = params.TeamID
	// }

	student.ID = id
	studentt, err := h.repo.GetStudentByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	user.ID = studentt.UserID
	_, errr := h.userrepo.UpdateUserByID(user.ID, user)
	if errr != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	updatedStudent, err := h.repo.UpdateStudentByID(id, student)
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
		Data: updatedStudent,
	})
}
func (h *StudentHandler) DeleteStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.repo.GetStudentByID(id)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	errrr := h.repo.DeleteStudentByID(id)
	if errrr != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	errr := h.userrepo.DeleteUserByID(student.UserID)
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
		Data: "",
	})
}

// Dependency Injection
func ProvideStudentHandler(repo StudentRepository, userrepo UserRepository) *StudentHandler {
	return &StudentHandler{
		repo:     repo,
		userrepo: userrepo,
	}
}
