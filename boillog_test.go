package boillog

import (
	"os"
	"strings"
	"testing"
	"time"
)

// TestEnvProfiler tests the envProfiler function.
func TestEnvProfiler(t *testing.T) {
	os.Setenv("PROFILER", "true")
	defer os.Unsetenv("PROFILER")

	expected := "true"
	if result := envProfiler(); result != expected {
		t.Errorf("Expected PROFILER to be %s, got %s", expected, result)
	}
}

// TestEnvLogLocation tests the envLogLocation function.
func TestEnvLogLocation(t *testing.T) {
	os.Setenv("LOGLOCATION", "/tmp/")
	defer os.Unsetenv("LOGLOCATION")

	expected := "/tmp/"
	if result := envLogLocation(); result != expected {
		t.Errorf("Expected LOGLOCATION to be %s, got %s", expected, result)
	}
}

// TestEnvAppName tests the envAppName function.
func TestEnvAppName(t *testing.T) {
	os.Setenv("APP_NAME", "testapp.log")
	defer os.Unsetenv("APP_NAME")

	expected := "testapp.log"
	if result := envAppName(); result != expected {
		t.Errorf("Expected APP_NAME to be %s, got %s", expected, result)
	}
}

// TestLogger tests the Logger function for successful log creation.
func TestLogger(t *testing.T) {
	testLogLocation := "/tmp/"
	testLogName := "test.log"

	// Set the environment variable for testing.
	os.Setenv("LOGLOCATION", testLogLocation)
	os.Setenv("APP_NAME", testLogName)
	defer func() {
		os.Unsetenv("LOGLOCATION")
		os.Unsetenv("APP_NAME")
		os.Remove(testLogLocation + testLogName) // Clean up after the test.
	}()

	if err := Logger("TestLogger", "INFO", "This is a test log message."); err != nil {
		t.Errorf("Logger should not have returned an error: %v", err)
	}

	// Read the log file and check contents.
	content, err := os.ReadFile(testLogLocation + testLogName)
	if err != nil {
		t.Fatalf("Error reading log file: %v", err)
	}

	if !strings.Contains(string(content), "This is a test log message.") {
		t.Errorf("Log file should contain the test log message.")
	}
}

// TestTrackTime tests the TrackTime function's profiling capability.
func TestTrackTime(t *testing.T) {
	startTime := time.Now()
	time.Sleep(10 * time.Millisecond)
	elapsed := TrackTime("TestTrackTime", startTime)

	if elapsed < 10*time.Millisecond {
		t.Errorf("TrackTime should have reported an elapsed time of at least 10ms, got %v", elapsed)
	}
}
