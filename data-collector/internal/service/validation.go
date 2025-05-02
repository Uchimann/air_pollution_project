package service

import(
	"time"
)

// validateTimestamp gelen Unix timestamp’in geçerli olup olmadığını kontrol eder.
// Geleceğe atılmışsa ErrInvalidTimestamp döner, değilse nil.
func validateTimestamp(ts int64) error {
    now := time.Now().Unix()
    if ts > now {
        return ErrInvalidTimestamp
    }
    return nil
}

func validatePollution(name string) error {
	
	if name == "PM2.5" || name == "PM10" || name == "NO2" || name == "SO2" || name == "O3"{
		return nil
	}
	return ErrUnsupportedPollutant
}