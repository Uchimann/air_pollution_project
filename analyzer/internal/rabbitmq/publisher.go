package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/uchimann/air_pollution_project/analyzer/internal/model"
)

func (c *Client) PublishAnalysisResult(result model.PollutionAnalysis) error {

	body, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error marshalling analysis result: %w", err)
	}

	err = c.ch.Publish(
		c.exchangeName,
		"notifications",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		
		return fmt.Errorf("error publishing message to RabbitMQ: %w", err)
	}	
	
	log.Printf("Analysis result published to RabbitMQ: %s (Level: %s)", 
        result.Pollutant, result.AnomalyLevel)
		
	return nil
}