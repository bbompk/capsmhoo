package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	user "capsmhoo/mono/user-service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	defer gracefulShutdown()

	initConfig()

	r := gin.Default()

	// Dependency Injection
	repo := user.ProvideRepository()
	handler := user.ProvideHandler(repo)

	// Init http endpoint routes
	user.ProvideRouter(r, handler)

	r.Run(":" + viper.GetString("user-service.port"))
}

// Read Config file
func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down server...")
}
