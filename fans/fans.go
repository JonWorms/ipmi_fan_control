package fans

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func ipmicommand(args ...string) (string, error) {
	out, err := exec.Command(ipmitool, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}

func GetNames() ([]string, error) {
	resp, err := ipmicommand("-c", "sdr")
	if err != nil {
		return nil, err
	}
	fans := []string{}

	lines := strings.Split(resp, "\n")
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		if strings.HasPrefix(tokens[0], "FAN") {
			fans = append(fans, tokens[0])
		}
	}

	return fans, nil
}

func SetManualControl(on bool) error {
	controlWord := "0x00"
	if !on {
		controlWord = "0x01"
	}
	_, err := ipmicommand("raw", "0x30", "0x30", "0x01", controlWord)
	return err
}

func SetFanSpeed(speed float32) error {

	controlWord := strconv.FormatInt(int64(speed), 16)
	if len(controlWord) == 1 {
		controlWord = fmt.Sprintf("0%s", controlWord)
	}

	controlWord = fmt.Sprintf("0x%s", controlWord)

	_, err := ipmicommand("raw", "0x30", "0x30", "0x02", "0xff", controlWord)
	return err
	//ipmitool -I lanplus -H $IP -U $USER -P $PASS raw 0x30 0x30 0x02 0xff 0x##
}
