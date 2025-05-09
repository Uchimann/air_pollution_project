package rabbitmq

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/uchimann/air_pollution_project/data-collector/internal/config"
)

type Client struct {
    conn         *amqp.Connection
    ch           *amqp.Channel
    exchangeName string
}

func NewClient() (*Client, error) {
	var RabbitMQHost = config.Get("RABBITMQ_HOST","rabbitmq")
	var RabbitMQPort = config.Get("RABBITMQ_PORT","5672")
	var RabbitMQUser = config.Get("RABBITMQ_USER","guest")
	var RabbitMQPassword = config.Get("RABBITMQ_PASSWORD","guest")
	
	url := "amqp://" + RabbitMQUser + ":" + RabbitMQPassword + "@" + RabbitMQHost + ":" + RabbitMQPort + "/"
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	var Client = &Client{
		conn:         conn,
		ch:           ch,
		exchangeName: "air_pollution_data",
	}

	err = ch.ExchangeDeclare(
		Client.exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	log.Println("RabbitMQ connection established")
	return Client, nil
}

func (c *Client) PublishPollutionData(data []byte) error {

    err := c.ch.Publish(
        c.exchangeName,
        "pollution",
        false,
        false,
        amqp.Publishing{
            ContentType:  "application/json",
            Body:         data,
            DeliveryMode: amqp.Persistent,
        })
    if err != nil {
        return fmt.Errorf("Mesaj yayınlanamadı: %w", err)
    }
    
    log.Println("Message published to RabbitMQ")
    return nil
}