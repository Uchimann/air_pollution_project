package rabbitmq

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/streadway/amqp"
    "github.com/uchimann/air_pollution_project/analyzer/internal/model"
)

func (c *Client) PublishAnomaly(analysis *model.PollutionAnalysis) error {
    data, err := json.Marshal(analysis)
    if err != nil {
        return fmt.Errorf("Analyze data cannot convert to JSON: %w", err)
    }

    err = c.ch.Publish(
        c.anomalyExchangeName,
        "anomaly", // routing key
        false,
        false,
        amqp.Publishing{
            ContentType:  "application/json",
            Body:         data,
            DeliveryMode: amqp.Persistent,
        },
    )
    if err != nil {
        return fmt.Errorf("Anomali message cannot published: %w", err)
    }

    log.Printf("Anomaly detected and published: %s - %f", analysis.Pollutant, analysis.Value)
    return nil
}