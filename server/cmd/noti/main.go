package main

import (
	"context"
	"fmt"
	"log"
	"os"

	noti "capsmhoo/mono/noti-service"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if os.Getenv("ENV") != "integration" && os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	initConfig()

	mongoClient, err := initMongoDB()
	if err != nil {
		panic(err)
	}

	// Dependency Injection
	repo := noti.ProvideNotiRepository(mongoClient)
	notiConsumer := noti.ProvideConsumer(repo)

	notis, err := repo.GetAllNotis()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(notis[0].Body)

	go notiConsumer.StartNotificationConsumer(viper.GetString("rabbitmq.url"), viper.GetString("rabbitmq.noti_queue_name"))
	noti.StartgRPCServer(repo, "", viper.GetString("noti-service.grpc-port"))
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

func initMongoDB() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return mongoClient, nil
}
