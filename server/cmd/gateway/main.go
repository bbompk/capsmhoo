package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	pb "capsmhoo/proto"
)

var db *gorm.DB

var addr = flag.String("addr", "localhost:8081", "127.0.0.1:8081")

type Team struct {
	Id      string `json:"id"`
	Name    string `json:"title"`
	Profile string `json:"profile"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTeamServiceClient(conn)
	log.Println("Client connected to grpc server...")

	r := gin.Default()

	r.GET("/team", func(c *gin.Context) {
		res, err := client.GetAllTeams(c, &pb.Empty{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var teams []*Team
		for _, team := range res.Teams {
			teams = append(teams, &Team{
				Id:      team.Id,
				Name:    team.Name,
				Profile: team.Profile,
			})
		}
		c.JSON(http.StatusOK, gin.H{"teams": teams})
	})
	r.GET("/team/:id", func(c *gin.Context) {
		id := c.Param("id")
		teamRes, err := client.GetTeamById(c, &pb.TeamId{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		team := &Team{
			Id:      teamRes.Id,
			Name:    teamRes.Name,
			Profile: teamRes.Profile,
		}
		c.JSON(http.StatusOK, gin.H{"team": team})
	})
	r.POST("/team", func(c *gin.Context) {
		var team Team
		if err := c.ShouldBindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		teamRes, err := client.CreateTeam(c, &pb.Team{Name: team.Name, Profile: team.Profile})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		team = Team{
			Id:      teamRes.Id,
			Name:    teamRes.Name,
			Profile: teamRes.Profile,
		}
		c.JSON(http.StatusOK, gin.H{"team": team})
	})
	r.PUT("/team/:id", func(c *gin.Context) {
		id := c.Param("id")
		var team Team
		if err := c.ShouldBindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		teamRes, err := client.UpdateTeam(c, &pb.Team{Id: id, Name: team.Name, Profile: team.Profile})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		team = Team{
			Id:      teamRes.Id,
			Name:    teamRes.Name,
			Profile: teamRes.Profile,
		}
		c.JSON(http.StatusOK, gin.H{"team": team})
	})
	r.DELETE("/team/:id", func(c *gin.Context) {
		id := c.Param("id")
		teamRes, err := client.DeleteTeam(c, &pb.TeamId{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		team := Team{
			Id:      teamRes.Id,
			Name:    teamRes.Name,
			Profile: teamRes.Profile,
		}
		c.JSON(http.StatusOK, gin.H{"team": team})
	})

	r.Run(":" + "8082")
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
