package analyzer

import (
	"log"
)

func AnomalyDetection(data []byte) (bool, error) {
	log.Printf("Anomaly detection started for data: %s\n", string(data))
	
	return true, nil
}
