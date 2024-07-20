package boillog

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

// Define a custom type to avoid collisions
type contextKey string

// Define constants for the key used
const (
	FuncKey contextKey = "func"
)

// LogIt Boilerplate funtion that calls Logger, to write logs, and prints it if it fails to write it
func LogIt(logFunction string, logOutput string, message string) {
	logPath := filepath.Join(envLogLocation(), envAppName())
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		// File doesn't exist, create it
		file, err := os.Create(logPath)
		if err != nil {
			log.Println("Error creating log file:", err)
			return
		}
		file.Close()
	}
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer file.Close()
	errCloseLogger := logger(logFunction, logOutput, message, file)
	if errCloseLogger != nil {
		log.Println(errCloseLogger)
	}
}

// Logger This function is called by Logit and prints/writes logs
func logger(logFunction string, logOutput string, message string, w io.Writer) error {
	timeNow := time.Now()
	handler := slog.NewTextHandler(w, nil)
	logger := slog.New(handler)
	ctx := context.Background()
	ctx = context.WithValue(ctx, slog.TimeKey, timeNow)
	ctx = context.WithValue(ctx, FuncKey, logFunction)
	switch logOutput {
	case "INFO":
		logger.InfoContext(ctx, message)
	case "WARNING":
		logger.WarnContext(ctx, message)
	case "ERROR":
		logger.ErrorContext(ctx, message)
	default:
		logger.InfoContext(ctx, message)
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
