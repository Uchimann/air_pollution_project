package rabbitmq

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/streadway/amqp"
    "github.com/uchimann/air_pollution_project/analyzer/internal/model"
    "github.com/uchimann/air_pollution_project/analyzer/internal/service"
)

type Consumer struct {
    client  *Client
    service *service.AnalyzerService
}

func NewConsumer(client *Client, service *service.AnalyzerService) (*Consumer, error) {
    consumer := &Consumer{
        client:  client,
        service: service,
    }

    queue, err := client.ch.QueueDeclare(
        client.queueName,
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return nil, fmt.Errorf("Queue cannot create: %w", err)
    }

    err = client.ch.QueueBind(
        queue.Name,
        "pollution", 
        client.exchangeName,
        false,
        nil,
    )
    if err != nil {
        return nil, fmt.Errorf("Queue and exchange cannot connect: %w", err)
    }

    return consumer, nil
}

func (c *Consumer) StartConsuming() error {
    msgs, err := c.client.ch.Consume(
        c.client.queueName,
        "",    
        false, 
        false, 
        false, 
        false, 
        nil,
    )
    if err != nil {
        return fmt.Errorf("Cunsomer cannot start: %w", err)
    }

    forever := make(chan bool)

    go func() {
        for msg := range msgs {
            log.Printf("New message: %s", msg.Body)
            
            var pollutantData model.PollutantData
            if err := json.Unmarshal(msg.Body, &pollutantData); err != nil {
                log.Printf("An error occured while read message: %v", err)
                msg.Nack(false, true) 
                continue
            }

            analysis, err := c.service.AnalyzePollutionData(&pollutantData)
            if err != nil {
                log.Printf("An error while data analyze: %v", err)
                msg.Nack(false, true)
                continue
            }

            if analysis.IsAnomalous {
                if err := c.service.NotifyAnomaly(analysis); err != nil {
                    log.Printf("Anomaly error: %v", err)
                }
            }

            msg.Ack(false)
        }
    }()

    log.Printf("RabbitMQ messages listening. Press CTRL+C'ye for exit")
    <-forever

    return nil
}