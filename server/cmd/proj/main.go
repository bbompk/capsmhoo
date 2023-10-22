package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	project "capsmhoo/mono/project-service"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	defer gracefulShutdown()

	initConfig()

	db, err := initDatabase()

	if err != nil {
		panic("Can't connect to Database")
	}

	// Dependency Injection
	repo := project.ProvideRepository(db)

	project.StartgRPCServer(repo, "", viper.GetString("project-service.grpc-port"))
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

// type Project struct {
// 	Id      string `json:"id"`
// 	Name    string `json:"title"`
// 	Profile string `json:"profile"`
// }

func initDatabase() (*gorm.DB, error) {
	// Read database connection parameters from config or environment variables
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbUser := viper.GetString("db.POSTGRES_USER")
	dbPassword := viper.GetString("db.POSTGRES_ROOT_PASSWORD")
	dbName := viper.GetString("db.POSTGRES_DB")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Println(connStr)
	// Open a database connection
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down server...")
}
