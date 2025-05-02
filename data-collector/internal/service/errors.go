
package service

import "errors"


var ErrInvalidTimestamp = errors.New("timestamp cannot be in the future")
var ErrUnsupportedPollutant = errors.New("Pollutant must be: 'PM2.5', 'PM10', 'NO2', 'SO2', 'O3' ")