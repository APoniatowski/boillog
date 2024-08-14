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

// TestEnvLogLevel tests the envLogLevel function.
func TestEnvLogLevel(t *testing.T) {
	testCases := []struct {
		envValue string
		expected string
	}{
		{"DEBUG", "DEBUG"},
		{"INFO", "INFO"},
		{"WARN", "WARN"},
		{"ERROR", "ERROR"},
		{"DISABLE", "DISABLE"},
		{"", "DEBUG"}, // Test for the default value when LOGLEVEL is not set
	}

	for _, tc := range testCases {
		if tc.envValue != "" {
			os.Setenv("LOGLEVEL", tc.envValue)
		} else {
			os.Unsetenv("LOGLEVEL")
		}

		result := envLogLevel()
		if result != tc.expected {
			t.Errorf("Expected LOGLEVEL to be %s, got %s", tc.expected, result)
		}

		os.Unsetenv("LOGLEVEL")
	}
}

// TestLogIt tests the LogIt function for successful log creation based on different log levels.
func TestLogIt(t *testing.T) {
	testLogLocation := "/tmp/"
	testLogName := "test.log"

	os.Setenv("LOGLOCATION", testLogLocation)
	os.Setenv("APP_NAME", testLogName)

	defer func() {
		os.Unsetenv("LOGLOCATION")
		os.Unsetenv("APP_NAME")
		os.Remove(testLogLocation + testLogName)
	}()

	testCases := []struct {
		logLevel      string
		logType       string
		message       string
		shouldContain bool
	}{
		{"DEBUG", "INFO", "This is an INFO log.", true},
		{"DEBUG", "DEBUG", "This is a DEBUG log.", true},
		{"INFO", "DEBUG", "This DEBUG log should not be logged.", false},
		{"INFO", "INFO", "This is another INFO log.", true},
		{"WARN", "INFO", "This INFO log should not be logged.", false},
		{"WARN", "WARN", "This is a WARN log.", true},
		{"ERROR", "WARN", "This WARN log should not be logged.", false},
		{"ERROR", "ERROR", "This is an ERROR log.", true},
		{"DISABLE", "ERROR", "This ERROR log should not be logged.", false},
	}

	for _, tc := range testCases {
		os.Setenv("LOGLEVEL", tc.logLevel)

		LogIt("TestLogIt", tc.logType, tc.message)

		content, err := os.ReadFile(testLogLocation + testLogName)
		if err != nil {
			t.Fatalf("Error reading log file: %v", err)
		}
		logContent := string(content)

		if tc.shouldContain {
			if !strings.Contains(logContent, tc.message) {
				t.Errorf("Expected log message not found. LogLevel: %s, LogType: %s, Message: %s", tc.logLevel, tc.logType, tc.message)
			}
		} else {
			if strings.Contains(logContent, tc.message) {
				t.Errorf("Unexpected log message found. LogLevel: %s, LogType: %s, Message: %s", tc.logLevel, tc.logType, tc.message)
			}
		}

		os.Remove(testLogLocation + testLogName)
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
