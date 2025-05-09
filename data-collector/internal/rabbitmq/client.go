package rabbitmq

import(
	"fmt"
	"log"
)

func (c *Client) Close() error {
    if c.ch != nil {
        if err := c.ch.Close(); err != nil {
            log.Printf("Error while closing channel: %v", err)
        }
    }
    
    if c.conn != nil {
        if err := c.conn.Close(); err != nil {
            return fmt.Errorf("Error while closing connection: %w", err)
        }
    }
    
    return nil
}