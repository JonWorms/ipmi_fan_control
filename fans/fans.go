package fans

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonworms/ipmi_fan_control/ipmi"
)

func GetNames() ([]string, error) {
	resp, err := ipmi.Command("-c", "sdr")
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
	_, err := ipmi.Command("raw", "0x30", "0x30", "0x01", controlWord)
	return err
}

func formatWord(value int64) string {
	word := strconv.FormatInt(value, 16)
	if len(word) == 1 {
		word = fmt.Sprintf("0%s", word)
	}
	return fmt.Sprintf("0x%s", word)
}

func SetFanSpeed(fanBitmask uint, speed float32) error {

	controlWord := formatWord(int64(speed))
	fanMaskWord := formatWord(int64(fanBitmask))

	_, err := ipmi.Command("raw", "0x30", "0x30", "0x02", fanMaskWord, controlWord)
	return err
	//ipmitool -I lanplus -H $IP -U $USER -P $PASS raw 0x30 0x30 0x02 0xff 0x##
}
