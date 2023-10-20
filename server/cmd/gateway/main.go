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
	gatewaygRPCClient "capsmhoo/mono/api-gateway/client_grpc"
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

	defer teamgRPCConn.Close()
	defer notigRPCConn.Close()

	// dependency injection
	teamgRPCClient := gatewaygRPCClient.ProvideTeamClient(&teamgRPCClienter)
	teamHandler := gatewayHTTPHandler.ProvideTeamHandler(teamgRPCClient)
	notigRPCClient := gatewaygRPCClient.ProvideNotiClient(&notigRPCClienter)
	notiHandler := gatewayHTTPHandler.ProvideNotiHandler(notigRPCClient)

	gatewayHTTPHandler.ProvideRouter(r, teamHandler, notiHandler)

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
