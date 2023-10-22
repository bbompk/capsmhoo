package middleware

import (
	"os"

	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuthentication(c *gin.Context) {
	// Get token from the 'Authorization' header.
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Check if the token format is correct
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token format is 'Bearer <token>'"})
		return
	}

	// Extract token from the split
	requestToken := splitToken[1]

	// Parse and validate the token
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY") 
		return []byte(jwtSecret), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		fmt.Printf("User ID: %v", claims["user_id"]) // "user_id" is one of the claims you put in the token
		c.Set("userID", claims["user_id"])

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
}