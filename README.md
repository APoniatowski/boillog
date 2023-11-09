# boillog

Before cooking, one has to boil something.

BoilLog is a Go package that provides a simple, yet flexible, logging library designed to integrate with Go applications. It supports environmental configuration and includes basic performance profiling features.

## Features

- **Environmental Configuration**: Customize log location, application name, and profiling preference through environment variables.
- **Boilerplate Logging**: Simplified logging interface to log messages with minimal setup.
- **Performance Profiling**: Optional performance profiling to measure and output the execution time of code blocks.

## Installation

To install BoilLog, use `go get` to retrieve the package:

```bash
go get github.com/APoniatowski/boillog
```

Make sure to import the package with:
```go
import "github.com/APoniatowski/boillog"
```

## Usage
### Setting Environment Variables

Set the following environment variables to configure BoilLog:

    PROFILER: Set to true to enable performance profiling (default: false).
    LOGLOCATION: The directory where log files will be stored (default: /var/log/).
    APP_NAME: The name of the log file (default: boiler.log).

## Logging Functions
### LogIt

Use LogIt to log messages easily. It takes a log function name, log type, and message as parameters:
```go
boillog.LogIt("MyFunction", "INFO", "This is an informative message.")
```

### Logger

Logger is used internally by LogIt but can also be used directly if desired. It performs the actual logging operation:
```go
err := boillog.Logger("MyFunction", "ERROR", "This is an error message.")
if err != nil {
    log.Println("Failed to log message:", err)
}
```

## Profiling

Use TrackTime to measure the execution time of functions for performance analysis:
```go
defer boillog.TrackTime("TaskName", time.Now())
// ... code to profile ...
```

# Contributing

Contributions to BoilLog are welcome. Please submit a pull request or raise an issue for any features or bugs.

# License

BoilLog is open-sourced software licensed under the MIT license.

