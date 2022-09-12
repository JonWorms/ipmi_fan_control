package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jonworms/ipmi_fan_control/fans"
	"github.com/jonworms/ipmi_fan_control/logging"
	"github.com/jonworms/ipmi_fan_control/sensors"
	"github.com/jonworms/ipmi_fan_control/speedmap"
	"github.com/soellman/pidfile"

	"github.com/sevlyar/go-daemon"
)

var mapPath string
var verboseLevel int
var pidfilePath string
var logfilePath string

func adjustFans() (float32, float32, error) {

	temp, err := sensors.GetAverageTemp()
	if err != nil {
		return 0, 0, err
	}

	fs := speedmap.GetFanSpeed(temp)

	fans.SetFanSpeed(fs)

	return temp, fs, nil
}

func loadMap() {

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
}

func daemonMain(context *daemon.Context) {
	defer context.Release()

	if logfilePath != "" {
		err := logging.Open(logfilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer logging.Close()
	}

	// create pid with child process pid
	if pidfilePath != "" {
		err := pidfile.Write(pidfilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer pidfile.Remove(pidfilePath)
	}

	go func() {
		err := fans.SetManualControl(true)
		if err != nil {
			logging.Fatal(err)
		}

		for {
			t, fs, err := adjustFans()
			if err != nil {
				logging.PrintError(err)
			}
			if verboseLevel > 0 {
				logging.LogTFS(t, fs)
			}

			time.Sleep(time.Duration(time.Second))
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-sigchan
}

func main() {

	flag.StringVar(&mapPath, "s", "/etc/fan_speeds.cfg", "Path to speed map config file")
	flag.IntVar(&verboseLevel, "v", 0, "verbosity level")
	flag.StringVar(&pidfilePath, "p", "", "path to pid file")
	flag.StringVar(&logfilePath, "l", "", "path to log file")
	flag.Parse()

	fans.SetIPMITool("/usr/local/bin/ipmitool")

	if verboseLevel > 0 {
		log.Println("starting in verbose mode")
	}

	loadMap()

	context := new(daemon.Context)
	child, _ := context.Reborn()
	if child == nil {
		daemonMain(context)
	}
}
