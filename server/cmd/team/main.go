package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	//"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	pb "capsmhoo/proto"
)

var db *gorm.DB

func main() {
	defer gracefulShutdown()

	initConfig()

	var err error
	db, err = initDatabase()

	if err != nil {
		panic("Can't connect to Database")
	}

	startServer()

	/*r := gin.Default()

	// Dependency Injection
	repo := team.ProvideRepository(db)
	handler := team.ProvideTeamHandler(repo)

	// Init http endpoint routes
	team.ProvideRouter(r, handler)

	r.Run(":" + viper.GetString("team-service.port"))*/
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

type Team struct {
	Id      string `json:"id"`
	Name    string `json:"title"`
	Profile string `json:"profile"`
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

var (
	port = flag.Int("port", 8081, "gRPC server port")
)

type teamServer struct {
	// Implements the generated TeamServer interface
	pb.UnimplementedTeamServiceServer
}

func (s *teamServer) mustEmbedUnimplementedTeamServiceServer() {}

func (s *teamServer) GetAllTeams(ctx context.Context, empty *pb.Empty) (*pb.TeamList, error) {
	fmt.Println("Get Teams")
	teams := []*pb.Team{}
	res := db.Find(&teams)
	if res.RowsAffected == 0 {
		return nil, errors.New("Team not found")
	}
	return &pb.TeamList{
		Teams: teams,
	}, nil
}

func (s *teamServer) GetTeamById(ctx context.Context, teamId *pb.TeamId) (*pb.Team, error) {
	fmt.Println("Get Team By ID")
	var team pb.Team
	res := db.Find(&team, "id = ?", teamId)
	if res.RowsAffected == 0 {
		return nil, errors.New("Team not found")
	}
	return &team, nil
}

func (s *teamServer) CreateTeam(ctx context.Context, team *pb.Team) (*pb.Team, error) {
	fmt.Println("Create Team")

	data := pb.Team{
		Id:      uuid.New().String(),
		Name:    team.Name,
		Profile: team.Profile,
	}

	res := db.Create(&data)
	if res.RowsAffected == 0 {
		return nil, errors.New("team creation unsuccessful")
	}

	return &data, nil
}

func (s *teamServer) UpdateTeam(ctx context.Context, team *pb.Team) (*pb.Team, error) {
	fmt.Println("Update Team")

	res := db.Model(&team).Where("id=?", team.Id).Updates(Team{Name: team.Name, Profile: team.Profile})

	if res.RowsAffected == 0 {
		return nil, errors.New("team not found")
	}

	return team, nil
}

func (s *teamServer) DeleteTeam(ctx context.Context, teamId *pb.TeamId) (*pb.Team, error) {
	fmt.Println("Delete Team")
	var team pb.Team
	res := db.Where("id = ?", teamId.Id).Delete(&team)
	if res.RowsAffected == 0 {
		return nil, errors.New("Team not found")
	}

	return &team, nil
}

func (s *teamServer) AddStudentToTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
	return nil, nil
}

func (s *teamServer) RemoveStudentFromTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
	return nil, nil
}

func startServer() {
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTeamServiceServer(s, &teamServer{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
	fmt.Println("Go gRPC server started")
}
