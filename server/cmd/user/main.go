package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	user "capsmhoo/internal/user-service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	defer gracefulShutdown()

	if os.Getenv("ENV") != "integration" && os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	initConfig()
	db, err := initDatabase()

	if err != nil {
		panic("Can't connect to Database")
	}
	r := gin.Default()

	// Dependency Injection
	repo := user.ProvideRepository(db)
	handler := user.ProvideUserHandler(repo)

	studentdb, studenterr := initDatabase()
	if studenterr != nil {
		panic("Can't connect to Database")
	}
	studentrepo := user.ProvideStudentRepository(studentdb)
	studenthandler := user.ProvideStudentHandler(studentrepo, repo)

	professordb, professorerr := initDatabase()
	if professorerr != nil {
		panic("Can't connect to Database")
	}
	professorrepo := user.ProvideProfessorRepository(professordb)
	professorhandler := user.ProvideProfessorHandler(professorrepo, repo)

	user.ProvideRouter(r, handler, studenthandler, professorhandler)
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
