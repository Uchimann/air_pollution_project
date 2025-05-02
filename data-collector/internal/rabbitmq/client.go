package rabbitmq

import(
	"fmt"
	"log"
)

func (c *Client) Close() error {
    if c.ch != nil {
        if err := c.ch.Close(); err != nil {
            log.Printf("Kanal kapatılırken hata oluştu: %v", err)
        }
    }
    
    if c.conn != nil {
        if err := c.conn.Close(); err != nil {
            return fmt.Errorf("Bağlantı kapatılırken hata oluştu: %w", err)
        }
    }
    
    return nil
}