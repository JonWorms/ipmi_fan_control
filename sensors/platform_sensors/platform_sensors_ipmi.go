package platformsensors

import (
	"errors"
	"strings"

	"github.com/jonworms/ipmi_fan_control/ipmi"
)

func GetIPMISensor(name string) (map[string]string, error) {
	rval, err := ipmi.Command("sensor", "get", name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(rval, "\n")
	values := make(map[string]string)
	for _, line := range lines {
		if strings.Contains(line, ":") {
			kvp := strings.Split(line, ":")
			values[strings.TrimSpace(kvp[0])] = strings.TrimSpace(kvp[1])
		}
	}
	return values, nil
}

func GetIPMISensorReading(name string) (string, error) {
	values, err := GetIPMISensor(name)
	if err != nil {
		return "", err
	}
	if value, ok := values["Sensor Reading"]; ok {
		return value, nil
	}
	return "", errors.New("Sensor Reading not found")
}
