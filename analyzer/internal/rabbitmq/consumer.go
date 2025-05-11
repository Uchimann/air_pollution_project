package rabbitmq

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func (c *Client) ConsumeMessages(q amqp.Queue) (chan []byte, error) {

	messageChan := make(chan []byte)
	
	msgs, err := c.ch.Consume(
		q.Name, // queue namei de dene q.name olarak
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %s", err)
	}

    go func() {
        for msg := range msgs {
            log.Printf("New message: %d bytes", len(msg.Body))
            messageChan <- msg.Body
        }
        log.Println("RabbitMQ message channel closed")
        close(messageChan)
    }()

	return messageChan, nil
}

func (c *Client) ConnectQueue() (amqp.Queue, error) {
    queueName := "analyzer_queue"
    
    q, err := c.ch.QueueDeclare(
        queueName,
        true,
        false,
        false,
        false,
        nil, 
    )
    if err != nil {
        return q, err 
    }
    

    err = c.ch.QueueBind(
        q.Name,
        "pollution",
        c.exchangeName,
        false,
        nil,
    )
    if err != nil {
        return q, err
    }

	log.Printf("Queue '%s' declared and bound to exchange '%s'", q.Name, c.exchangeName)

	return q,nil
}