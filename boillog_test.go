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

// TestLogIt tests the LogIt function for successful log creation.
func TestLogIt(t *testing.T) {
	testLogLocation := "/tmp/"
	testLogName := "test.log"
	os.Setenv("LOGLOCATION", testLogLocation)
	os.Setenv("APP_NAME", testLogName)
	defer func() {
		os.Unsetenv("LOGLOCATION")
		os.Unsetenv("APP_NAME")
		os.Remove(testLogLocation + testLogName) // Clean up after the test.
	}()
	LogIt("TestLogIt", "INFO", "This is a test log message.")
	content, err := os.ReadFile(testLogLocation + testLogName)
	if err != nil {
		t.Fatalf("Error reading log file: %v", err)
	}
	logContent := string(content)
	if !strings.Contains(logContent, "This is a test log message.") {
		t.Errorf("Log file should contain the test log message.")
	}
}

// TestTrackTime tests the TrackTime function's profiling capability.
func TestTrackTime(t *testing.T) {
	startTime := time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := TrackTime("TestTrackTime", startTime)
	if elapsed < 100*time.Millisecond {
		t.Errorf("TrackTime should have reported an elapsed time of at least 10ms, got %v", elapsed)
	}
}
