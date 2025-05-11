package rabbitmq

import (
    "github.com/streadway/amqp"
    "log"
)

type Client struct {
    conn         *amqp.Connection
    ch           *amqp.Channel
    exchangeName string
}

func NewClient() (*Client, error) {
    conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
    if err != nil {
        return nil, err
    }
    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }
    client := &Client{
        conn:         conn,
        ch:           ch,
        exchangeName: "air_pollution_data",
    }
    err = ch.ExchangeDeclare(
        client.exchangeName, "direct", true, false, false, false, nil,
    )
    if err != nil {
        return nil, err
    }
    log.Println("RabbitMQ successfully connected")
    return client, nil
}

func (c *Client) Close() {
    c.ch.Close()
    c.conn.Close()
}