package service

import "errors"

var ErrUnsupportedPollutant = errors.New("unspported pollutant")
var ErrAnalysisFailed = errors.New("analysis failed")