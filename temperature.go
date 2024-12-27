package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func getWaterTemp() (float64, error) {
	var request, httpErr = http.Get(fmt.Sprintf("http://%s/homekit/temperature", localIp))
	if httpErr != nil {
		return 0, httpErr
	}
	var temperatureBytes, _ = io.ReadAll(request.Body)

	var temperatureAsFloat, parseError = strconv.ParseFloat(string(temperatureBytes), 16)

	if parseError != nil {
		return 0, parseError
	}

	return temperatureAsFloat, nil
}
