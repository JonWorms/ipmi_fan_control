//go:build freebsd || darwin

package platformsensors

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func parseCPUTemp(reading string) (float32, error) {
	tokens := strings.Split(reading, " ")
	for _, token := range tokens {
		if strings.HasSuffix(token, "C") {
			token = token[0 : len(token)-1]
			val, err := strconv.ParseFloat(token, 32)
			return float32(val), err
		}
	}
	return 0, errors.New("parse error")
}

func GetCPUTemps() ([]float32, error) {
	out, err := exec.Command("sysctl", "dev.cpu").Output()
	if err != nil {
		return nil, err
	}

	temps := []float32{}

	lines := strings.Split(string(out[:]), "\n")
	for _, line := range lines {
		if strings.Contains(line, "temperature") {
			temp, err := parseCPUTemp(line)
			if err != nil {
				return nil, err
			}
			temps = append(temps, temp)
		}
	}

	return temps, nil
}
