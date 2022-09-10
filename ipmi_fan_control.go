package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jonworms/ipmi_fan_control/fans"
	"github.com/jonworms/ipmi_fan_control/sensors"
	"github.com/jonworms/ipmi_fan_control/speedmap"
)

var mapPath string
var verboseLevel int

type speedMapPoint struct {
	Temperature float32 `json:"temperature"`
	Speed       float32 `json:"speed"`
}

func adjustFans() error {
	temp, err := sensors.GetAverageTemp()

	if err != nil {
		return err
	}

	fs := speedmap.GetFanSpeed(temp)
	if verboseLevel > 0 {
		log.Printf("temp: %0.2f, speed: %0.2f", temp, fs)
	}
	fans.SetFanSpeed(fs)

	return nil
}

func main() {

	flag.StringVar(&mapPath, "s", "/etc/fan_speeds.cfg", "Path to speed map config file")
	flag.IntVar(&verboseLevel, "v", 0, "verbosity level")
	flag.Parse()

	if verboseLevel > 0 {
		log.Println("starting in verbose mode")
	}

	log.Printf("attempting to load map from %s\n", mapPath)

	err := speedmap.LoadFromFile(mapPath)
	if err != nil {
		log.Println(err)
	}
	message := "speed map:"
	for _, point := range speedmap.Points() {
		message = fmt.Sprintf("%s\n    %s", message, point.String())
	}
	log.Println(message)

	fans.SetManualControl(true)
	for true {
		err = adjustFans()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(time.Second))
	}
}
