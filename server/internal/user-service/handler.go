package user

import (
	"capsmhoo/common"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Define Dependencies
type UserHandler struct {
	repo UserRepository
}

// Define what this will do
type UserHttpHandler interface {
	// SignUpUser(c *gin.Context)
	SignInUser(c *gin.Context)
	SignOutUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUserByID(c *gin.Context)
	DeleteUserByID(c *gin.Context)
	DeleteAll(c *gin.Context)
}

// func (h* UserHandler) SignUpUser(c *gin.Context) {
// 	var user User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
// 		return
// 	}
// 	user.Password = string(hashedPassword)

// 	// Call the repository to save the user
// 	createdUser, err := h.repo.CreateUser(user)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
// }

func (h *UserHandler) SignInUser(c *gin.Context) {
	var userInput User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user by email or username
	user, err := h.repo.GetUserByEmail(userInput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	// Compare the provided password with the hash stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT token for the authenticated user
	token, err := CreateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func CreateToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expiration 1 hour from now

	jwtSecret := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	tokenString, err := token.SignedString([]byte(jwtSecret)) // <- Secret key (keep this safe!)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *UserHandler) SignOutUser(c *gin.Context) {
	// todo
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user := h.repo.GetUsers()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: user,
	})
}
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.repo.GetUserByID(id)
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
		Data: user,
	})
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
	// 	return
	// }
	// user.Password = string(hashedPassword)
	createdUser, err := h.repo.CreateUser(user)
	if err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	// user := h.repo.CreateUser()

	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: createdUser,
	})
}
func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, common.HttpResponse{
			Code: "400",
			// Data: {},
			Error: err.Error(),
		})
		return
	}
	updatedUser, err := h.repo.UpdateUserByID(id, user)
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
		Data: updatedUser,
	})
}
func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.DeleteUserByID(id)
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
		Data: User{},
	})
}
func (h *UserHandler) DeleteAll(c *gin.Context) {
	h.repo.DeleteAll()
	c.JSON(200, common.HttpResponse{
		Code: "200",
		Data: "",
	})
}

// Dependency Injection
func ProvideUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}
