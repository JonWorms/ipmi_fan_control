package sensors

import (
	"math"

	platformsensors "github.com/jonworms/ipmi_fan_control/sensors/platform_sensors"
)

func GetHottestTemp() (float32, error) {
	temps, err := platformsensors.GetCPUTemps()
	if err != nil {
		return 0.0, err
	}

	hottestTemp := temps[0]
	for i := 1; i < len(temps); i++ {
		hottestTemp = float32(math.Max(float64(hottestTemp), float64(temps[i])))
	}

	return hottestTemp, nil
}

func GetAverageTemp() (float32, error) {
	temps, err := platformsensors.GetCPUTemps()
	if err != nil {
		return 0.0, err
	}
	averageTemp := float32(0.0)
	for _, temp := range temps {
		averageTemp += temp
	}
	return averageTemp / float32(len(temps)), nil
}
