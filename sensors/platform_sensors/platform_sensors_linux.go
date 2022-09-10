//go:build linux

package platformsensors

import (
	"strconv"
	"strings"

	"github.com/ssimunic/gosensors"
)

func parseCPUTemp(reading string) (float32, error) {
	tokens := strings.Split(reading, " ")
	tempReading := tokens[0]
	tempReading = tempReading[1 : len(tempReading)-3]

	temp, err := strconv.ParseFloat(tempReading, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(temp), nil
}

func GetCPUTemps() ([]float32, error) {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		return nil, err
	}

	readings := []float32{}

	for chip := range sensors.Chips {
		// Iterate over entries
		for key, value := range sensors.Chips[chip] {
			// If CPU or GPU, print out
			if strings.HasPrefix(key, "Core") {
				reading, err := parseCPUTemp(value)
				if err != nil {
					return nil, err
				}
				readings = append(readings, reading)
			}
		}
	}

	return readings, nil

}
