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
    queueName    string
    exchangeName string
    anomalyExchangeName string
}

func NewClient() (*Client, error) {
    var RabbitMQHost = config.Get("RABBITMQ_HOST", "rabbitmq")
    var RabbitMQPort = config.Get("RABBITMQ_PORT", "5672")
    var RabbitMQUser = config.Get("RABBITMQ_USER", "guest")
    var RabbitMQPassword = config.Get("RABBITMQ_PASSWORD", "guest")

    url := "amqp://" + RabbitMQUser + ":" + RabbitMQPassword + "@" + RabbitMQHost + ":" + RabbitMQPort + "/"
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, fmt.Errorf("RabbitMQ connection error: %w", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, fmt.Errorf("RabbitMQ channel error: %w", err)
    }

    client := &Client{
        conn:         conn,
        ch:           ch,
        exchangeName: "air_pollution_data",
        queueName:    "analyzer_queue",
        anomalyExchangeName: "air_pollution_anomalies",
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
        client.Close()
        return nil, fmt.Errorf("Data collector exchange error: %w", err)
    }

    err = ch.ExchangeDeclare(
        client.anomalyExchangeName,
        "direct",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        client.Close()
        return nil, fmt.Errorf("Anomali exchange error: %w", err)
    }

    log.Println("RabbitMQ connected successfully")
    return client, nil
}

func (c *Client) Close() error {
    var err error
    
    if c.ch != nil {
        if err = c.ch.Close(); err != nil {
            log.Printf("An error occured while closing channel %v", err)
        }
    }

    if c.conn != nil {
        if err = c.conn.Close(); err != nil {
            return fmt.Errorf("An error occured while closing connection: %w", err)
        }
    }

    log.Println("RabbitMQ connection closed")
    return nil
}