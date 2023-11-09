package boillog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// ///////////////////////////== Environment Variables ==////////////////////////////
//
// envProfiler Gets the PROFILER environment variable (bool) and returns it
func envProfiler() string {
	profiler := os.Getenv("PROFILER")
	if len(profiler) == 0 {
		profiler = "false"
	}
	return profiler
}

// envLogLocation Gets the LOGLOCATION environment variable (bool) and returns it
func envLogLocation() string {
	logLocation := os.Getenv("LOGLOCATION")
	if len(logLocation) == 0 {
		logLocation = "/var/log/"
	}
	return logLocation
}

// envAppName Gets the APP_NAME environment variable (bool) and returns it
func envAppName() string {
	logName := os.Getenv("APP_NAME")
	if len(logName) == 0 {
		logName = "boiler.log"
	}
	return logName
}

///////////////////////////////== Logging functions ==//////////////////////////////

// LogIt Boilerplate funtion that calls Logger, to write/prints logs
func LogIt(logFunction string, logOutput string, message string) {
	errCloseLogger := Logger(logFunction, logOutput, message)
	if errCloseLogger != nil {
		log.Println(errCloseLogger)
	}
}

// Logger This function is called by Logit and prints/writes logs
func Logger(logFunction string, logOutput string, message string) error {
	currentDate := time.Now().Format("2006-01-02 15:04:05")
	pathString := envLogLocation()
	logName := envAppName()
	path, _ := filepath.Abs(pathString)
	err := os.MkdirAll(path, os.ModePerm)
	if err == nil || os.IsExist(err) {
		logFile, err := os.OpenFile(pathString+logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer func() {
			errClose := logFile.Close()
			if errClose != nil {
				log.Println(errClose)
			}
		}()
		logger := log.New(logFile, "", log.LstdFlags)
		logger.SetPrefix(currentDate)
		logger.Print(logFunction + " [ " + logOutput + " ] ==> " + message)
	} else {
		return err
	}
	if logOutput != "INFO" {
		fmt.Println("\t" + logFunction + " [ " + logOutput + " ] ==> " + message)
	}
	return nil
}

// TrackTime defer this function right at the beginning, to track time from start to finish
func TrackTime(taskName string, pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	profiler, err := strconv.ParseBool(envProfiler())
	if err != nil {
		fmt.Println(err)
	}
	if profiler {
		fmt.Printf("%v ", taskName)
		fmt.Println("elapsed:", elapsed)
	}
	return elapsed
}
