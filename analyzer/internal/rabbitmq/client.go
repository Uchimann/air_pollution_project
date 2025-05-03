package rabbitmq

import (
    "fmt"
    "log"
    
    "github.com/streadway/amqp"
    "github.com/uchimann/air_pollution_project/analyzer/internal/config"
)

type Client struct {
    conn         *amqp.Connection
    ch           *amqp.Channel
    exchangeName string
}

func NewClient() (*Client, error) {
    var RabbitMQHost = config.Get("RABBITMQ_HOST", "rabbitmq")
    var RabbitMQPort = config.Get("RABBITMQ_PORT", "5672")
    var RabbitMQUser = config.Get("RABBITMQ_USER", "guest")
    var RabbitMQPassword = config.Get("RABBITMQ_PASSWORD", "guest")
    
    url := "amqp://" + RabbitMQUser + ":" + RabbitMQPassword + "@" + RabbitMQHost + ":" + RabbitMQPort + "/"
    conn, err := amqp.Dial(url)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %s", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %s", err)
    }
    
    var client = &Client{
        conn:         conn,
        ch:           ch,
        exchangeName: "air_pollution_data",
    }

    err = ch.ExchangeDeclare(
        client.exchangeName,
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
    return client, nil
}

func (c *Client) Close() error {
    if c.ch != nil {
        if err := c.ch.Close(); err != nil {
            log.Printf("Error while close channel: %v", err)
        }
    }
    
    if c.conn != nil {
        if err := c.conn.Close(); err != nil {
            return fmt.Errorf("error while closing connection: %w", err)
        }
    }
    
    return nil
}