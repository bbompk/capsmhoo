package noti

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	repo NotiRepository
}

func (c *Consumer) StartNotificationConsumer(rabbitMQURL string, queueName string) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	} else {
		log.Printf("Connected to RabbitMQ: %s", rabbitMQURL)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue
	_, err = ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Start consuming messages from the queue: %s", queueName)

	// Process incoming messages
	for msg := range msgs {
		log.Printf("Received a noti to create: %s", msg.Body)

		var noti CreateNotificationMessage
		err := json.Unmarshal(msg.Body, &noti)
		if err != nil {
			log.Fatalf("Failed to unmarshal message: %v", err)
		}

		// Create a noti
		_, err = c.repo.CreateNoti(Notification{
			Title:  noti.Title,
			Body:   noti.Body,
			UserID: noti.ReceiverID,
			IsRead: false,
		})
		if err != nil {
			log.Fatalf("Failed to create a noti: %v", err)
		}
	}
}

func ProvideConsumer(repo NotiRepository) *Consumer {
	return &Consumer{repo: repo}
}
