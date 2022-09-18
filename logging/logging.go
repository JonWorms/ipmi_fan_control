package logging

import (
	"log"
	"os"
)

var logfile *os.File = nil
var logger *log.Logger = nil

const formatTFS = "temp: %0.2f, fan_speed: %0.2f\n"

var LogTFS func(float32, float32) = func(t, fs float32) {
	log.Printf(formatTFS, t, fs)
}

var Fatal func(error) = func(e error) {
	log.Fatal(e)
}

var PrintError = func(e error) {
	log.Println(e)
}

var Println = func(line string) {
	log.Println(line)
}

func Open(logPath string) error {
	log.Println("opening log file", logPath)
	logfile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logger = log.New(logfile, "", log.LstdFlags)

	LogTFS = func(t, fs float32) {
		logger.Printf(formatTFS, t, fs)
	}

	Fatal = func(e error) {
		logger.Fatal(e)
	}

	PrintError = func(e error) {
		logger.Println(e)
	}

	Println = func(line string) {
		logger.Println(line)
	}

	return nil
}

func Close() {
	logger = nil
	logfile.Close()
}
