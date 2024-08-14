package boillog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slog"
)

// ///////////////////////////== Environment Variables ==////////////////////////////

// envProfiler Gets the PROFILER environment variable (bool) and returns it
func envMetrics() string {
	profiler := os.Getenv("METRICS")

	if len(profiler) == 0 {
		profiler = "false"
	}

	return profiler
}

// envProfiler Gets the PROFILER environment variable (bool) and returns it
func envProfiler() string {
	profiler := os.Getenv("PROFILER")

	if len(profiler) == 0 {
		profiler = "false"
	}

	return profiler
}

// envLogLocation Gets the LOGLOCATION environment variable (string) and returns it
func envLogLocation() string {
	logLocation := os.Getenv("LOGLOCATION")

	if len(logLocation) == 0 {
		logLocation = "/var/log/"
	}

	return logLocation
}

// envAppName Gets the APP_NAME environment variable (string) and returns it
func envAppName() string {
	logName := os.Getenv("APP_NAME")

	if len(logName) == 0 {
		logName = "boiler.log"
	}

	return logName
}

// envLogLevel Gets the LOGLEVEL environment variable (string) and returns it
func envLogLevel() string {
	logName := os.Getenv("LOGLEVEL")

	if len(logName) == 0 {
		logName = "DEBUG"
	}

	return logName
}

///////////////////////////////== Logging functions ==//////////////////////////////

// Define a custom type to avoid collisions
type contextKey string

// Define constants for the key used
const (
	FuncKey contextKey = "func"
)

// LogIt Boilerplate funtion that calls Logger, to write logs, and prints it if it fails to write it
func LogIt(logFunction string, logType string, message string) {
	logPath := filepath.Join(envLogLocation(), envAppName())

	dir := filepath.Dir(logPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("Error creating log directory: %v", err)
		return
	}

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening or creating log file: %v", err)
		return
	}
	defer file.Close()

	errCloseLogger := logger(logFunction, strings.ToUpper(logType), message, file)
	if errCloseLogger != nil {
		log.Printf("Error writing to log: %v", errCloseLogger)
	}
}

// Logger This function is called by LogIt and prints/writes logs
func logger(logFunction string, logType string, message string, w io.Writer) error {
	logLevel := strings.ToUpper(envLogLevel())

	logLevels := map[string]int{
		"DEBUG":   0,
		"INFO":    1,
		"WARN":    2,
		"ERROR":   3,
		"DISABLE": 4,
	}

	currentLogLevel := logLevels[logLevel]
	messageLogLevel := logLevels[logType]

	if messageLogLevel < currentLogLevel || logLevel == "DISABLE" {
		return nil
	}

	// Continue with the logging process
	timeNow := time.Now().Format("2006-01-02 15:04:05") // Custom timestamp format

	handler := slog.NewTextHandler(w, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				return slog.Attr{Key: slog.TimeKey, Value: slog.StringValue(timeNow)}
			case slog.MessageKey:
				return a
			case slog.LevelKey:
				return slog.Attr{}
			default:
				return a
			}
		},
	})

	logger := slog.New(handler)

	logMessage := fmt.Sprintf("[%s] [%s] %s", timeNow, logFunction, message)

	switch logType {
	case "INFO":
		logger.Info(logMessage)
	case "WARN":
		logger.Warn(logMessage)
	case "ERROR":
		logger.Error(logMessage)
	default:
		logger.Info(logMessage)
	}

	return nil
}

// TrackTime defer this function right at the beginning, to track time from start to finish. And if you set METRICS to true, you'll get memory usage as well
func TrackTime(taskName string, pre time.Time) time.Duration {
	startCPU := new(runtime.MemStats)

	runtime.ReadMemStats(startCPU)

	profiler, err := strconv.ParseBool(envProfiler())
	if err != nil {
		fmt.Println(err)
	}

	metrics, err := strconv.ParseBool(envMetrics())
	if err != nil {
		fmt.Println(err)
	}

	elapsedCPU := new(runtime.MemStats)

	runtime.ReadMemStats(elapsedCPU)

	elapsed := time.Since(pre)
	if profiler {
		fmt.Printf("%v ", taskName)
		fmt.Println("elapsed:", elapsed)
	}

	if metrics {
		fmt.Println("Memory usage:", elapsedCPU.TotalAlloc-startCPU.TotalAlloc)
	}

	return elapsed
}
