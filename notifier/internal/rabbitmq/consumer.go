package rabbitmq

import (
    "encoding/json"
    "log"
	
    "github.com/uchimann/air_pollution_project/notifier/internal/model"
	"github.com/uchimann/air_pollution_project/notifier/internal/sse"
)

func (c *Client) ConsumeNotifications(es *sse.EventServer) error {
    q, err := c.ch.QueueDeclare(
        "notifier_queue", true, false, false, false, nil,
    )
    if err != nil {
        return err
    }
    err = c.ch.QueueBind(
        q.Name, "notifications", c.exchangeName, false, nil,
    )
    if err != nil {
        return err
    }
    msgs, err := c.ch.Consume(
        q.Name, "", true, false, false, false, nil,
    )
    if err != nil {
        return err
    }
    go func() {
        for msg := range msgs {
            var analysis model.PollutionAnalysis
            if err := json.Unmarshal(msg.Body, &analysis); err != nil {
                log.Printf("Error unmarshalling: %v", err)
                continue
            }
            es.Broadcast(analysis) // Sabit olarak sadece Broadcast çağrılır
        }
    }()
    return nil
}