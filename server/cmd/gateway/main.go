package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "capsmhoo/gen/proto"
	joinRequestPb "capsmhoo/gen/team-join-request-pb"
	gatewaygRPCClient "capsmhoo/mono/api-gateway/client_grpc"
	restClient "capsmhoo/mono/api-gateway/client_rest"
	gatewayHTTPHandler "capsmhoo/mono/api-gateway/http_handler"
)

type Team struct {
	Id      string `json:"id"`
	Name    string `json:"title"`
	Profile string `json:"profile"`
}

func main() {
	flag.Parse()

	defer gracefulShutdown()

	initConfig()

	r := gin.Default()
	teamgRPCConn := initTeamgRPCConnection()
	teamgRPCClienter := pb.NewTeamServiceClient(teamgRPCConn)
	notigRPCConn := initNotigRPCConnection()
	notigRPCClienter := pb.NewNotiServiceClient(notigRPCConn)
	teamJoinRequestgRPCClienter := joinRequestPb.NewTeamJoinRequestServiceClient(teamgRPCConn)

	defer teamgRPCConn.Close()
	defer notigRPCConn.Close()

	// dependency injection
	teamgRPCClient := gatewaygRPCClient.ProvideTeamClient(&teamgRPCClienter)
	teamJoinRequestgRPCClient := gatewaygRPCClient.ProvideTeamJoinRequestClient(&teamJoinRequestgRPCClienter)
	teamHandler := gatewayHTTPHandler.ProvideTeamHandler(teamgRPCClient)
	teamJoinRequestHandler := gatewayHTTPHandler.ProvideTeamJoinRequestHandler(teamJoinRequestgRPCClient)
	notigRPCClient := gatewaygRPCClient.ProvideNotiClient(&notigRPCClienter)
	notiHandler := gatewayHTTPHandler.ProvideNotiHandler(notigRPCClient)
	studentClientRest := restClient.ProvideStudentClientRest()
	studentHandler := gatewayHTTPHandler.ProvideStudentHandler(studentClientRest)
	professorClientRest := restClient.ProvideProfessorClientRest()
	professorHandler := gatewayHTTPHandler.ProvideProfessorHandler(professorClientRest)
	userClientRest := restClient.ProvideUserClientRest()
	userHandler := gatewayHTTPHandler.ProvideUserHandler(userClientRest)
	projectHandler := gatewayHTTPHandler.ProvideProjectHandler()

	gatewayHTTPHandler.ProvideRouter(r, teamHandler, teamJoinRequestHandler, userHandler, studentHandler, professorHandler, projectHandler, notiHandler)

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

func initTeamgRPCConnection() *grpc.ClientConn {
	dest := fmt.Sprintf("%s:%s", viper.GetString("team-service.grpc-host"), viper.GetString("team-service.grpc-port"))
	// Set up a connection to the server.
	conn, err := grpc.Dial(dest, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Default().Println("Connected to Team gRPC Service")
	return conn
}

func initNotigRPCConnection() *grpc.ClientConn {
	dest := fmt.Sprintf("%s:%s", viper.GetString("noti-service.grpc-host"), viper.GetString("noti-service.grpc-port"))
	// Set up a connection to the server.
	conn, err := grpc.Dial(dest, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Default().Println("Connected to Noti gRPC Service")
	return conn
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down server...")
}
