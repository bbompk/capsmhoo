package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

type RabbitMQPublisher interface {
	Close() error
	PublishMessage(ctx context.Context, exchange string, routingKey string, message interface{}) error
	PublishDefaultExchange(ctx context.Context, queue_name string, message interface{}) error
}

func (r *Publisher) Close() error {
	err := r.Channel.Close()
	if err != nil {
		log.Printf("Error closing rabbitmq channel: %v", err)
		return err
	}
	err = r.Conn.Close()
	if err != nil {
		log.Printf("Error closing rabbitmq connection %v", err)
		return err
	}
	return nil
}

func (r *Publisher) PublishMessage(ctx context.Context, exchange string, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.Channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *Publisher) PublishDefaultExchange(ctx context.Context, queue_name string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.Channel.PublishWithContext(
		ctx,
		"",
		queue_name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func ProvidePublisher(PublisherUrl string) (*Publisher, error) {
	conn, err := amqp091.Dial(PublisherUrl)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	log.Default().Println("RabbitMQ Publisher Connected")

	return &Publisher{
		Conn:    conn,
		Channel: ch,
	}, nil
}
