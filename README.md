# boillog

Before cooking, one has to boil something.

BoilLog is a Go package that provides a simple, yet flexible, logging library designed to integrate with Go applications. It supports environmental configuration and includes basic performance profiling and metrics features.

## Features

- **Environmental Configuration**: Customize log location, application name, profiling preference, log level output, and metrics output through environment variables.
- **Boilerplate Logging**: Simplified logging interface to log messages with minimal setup.
- **Performance Profiling**: Optional performance profiling to measure and output the execution time of code blocks.
- **Memory Usage Metrics**: Optional memory usage tracking for performance analysis.

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

- `PROFILER`: Set to `true` to enable performance profiling (default: `false`).
- `METRICS`: Set to `true` to enable memory usage metrics (default: `false`).
- `LOGLOCATION`: The directory where log files will be stored (default: `/var/log/`).
- `APP_NAME`: The name of the log file (default: `boiler.log`).

## Logging Functions

### LogIt

Use LogIt to log messages easily. It takes a log function name, log type, and message as parameters:

```go
boillog.LogIt("MyFunction", "INFO", "This is an informative message.")
```

## Profiling and Metrics

Use TrackTime to measure the execution time of functions and optionally track memory usage:

```go
defer boillog.TrackTime("TaskName", time.Now())
// ... code to profile ...
```

When `PROFILER` is set to `true`, this function will output the elapsed time for the task.
When `METRICS` is set to `true`, it will also output the memory usage for the task.

## Environment Variables

- `METRICS`: Controls whether memory usage metrics are reported (default: `false`).
- `PROFILER`: Controls whether execution time is reported (default: `false`).
- `LOGLOCATION`: Sets the directory for log files (default: `/var/log/`).
- `APP_NAME`: Sets the name of the log file (default: `boiler.log`).

## Contributing

Contributions to BoilLog are welcome. Please submit a pull request or raise an issue for any features or bugs.

## License

BoilLog is open-sourced software licensed under the MIT license.
