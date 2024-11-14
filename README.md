# Logma - A minimalistic Go logger with ✨Colours✨

Logma is a simple, lightweight logging library for Go that emphasizes simplicity as its strength. In many projects, more complex logging libraries are overkill and add unnecessary complexity. Logma solves this by providing an easy way to log messages with different levels of severity and ANSI colour-coding for improved readability in the terminal. Its minimalistic design allows you to focus on your core logic without the overhead of configuring complex logging systems.

In projects where logging needs are simple, you would otherwise end up writing a custom logger for each project. Logma gives you a reusable, straightforward solution, freeing you from reinventing the wheel every time.

## Table of contents

<!--toc:start-->

- [Logma - A minimalistic Go logger with ✨Colours✨](#logma-a-minimalistic-go-logger-with-colours)
  - [Table of contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Logger](#creating-a-logger)
    - [Logging Levels](#logging-levels)
    - [Example](#example)
    - [Output](#output)
    - [Log Level Behavior](#log-level-behavior)
    - [Colour Coding](#colour-coding)
    - [Exiting on Fatal Errors](#exiting-on-fatal-errors)
  - [License](#license)
  <!--toc:end-->

## Features

- **Logging Levels**: Control what levels of logs are displayed with inclusive logging:
  - `Fatal`: Logs critical errors that exit the program.
  - `Error`: Logs errors and `Fatal` messages.
  - `Info`: Logs informational, `Error`, and `Fatal` messages.
  - `All`: Logs all levels, including debug messages.
- **Colour Output**: Colour-coded log levels to make logs easier to read:
  - **Info**: Bold Green
  - **Debug**: Bold Yellow
  - **Error**: Bold Red
  - **Fatal**: Bold Yellow on Red Background

## Installation

Install `logma` using:

```bash
go get github.com/NewDawn0/logma
```

## Usage

To get started, import the `logma` package and create a new `Logger` instance. You can specify the log level to control the verbosity of the output.

### Creating a Logger

There are two ways to create a `Logger`:

1. **Custom Logger with Specific Log Level**:

   ```go
   logger := logma.NewLogger(logma.Info)
   ```

2. **Default Logger (Error level)**:
   ```go
   logger := logma.DefaultLogger()
   ```

### Logging Levels

The logging levels are defined as constants in `logma`. Each level is **inclusive** of all higher-severity levels:

- `logma.Fatal`: Solely fatal messages.
- `logma.Error`: Includes error and fatal messages.
- `logma.Info`: Includes info, error, and fatal messages.
- `logma.All`: Includes all messages, including debug.

### Example

```go
package main

import (
    "github.com/NewDawn0/logma"
)

func main() {
    // Create a logger with Info level
    logger := logma.NewLogger(logma.Info)

    // Log messages with various levels
    logger.Info("This is an informational message: %s", "Logma is active!")
    logger.Dbg("This is a debug message: %d items processed", 42)  // Won't show if log level < All
    logger.Error("An error occurred: %v", "file not found")
    logger.Fatal("Fatal error, terminating program: %s", "unable to connect to server")
}
```

### Output

With `logma.Info` as the log level, the output might look like this:

```plaintext
[Info] This is an informational message: Logma is active!
[Error] An error occurred: file not found
[Fatal] Fatal error, terminating program: unable to connect to server
```

### Log Level Behavior

Each log level in `Logma` includes all subsequent levels:

- **Fatal**: Logs exclusively fatal messages.
- **Error**: Logs both error and fatal messages.
- **Info**: Logs info, error, and fatal messages.
- **All**: Logs all messages, including debug.

### Colour Coding

Each log level is colour-coded for readability:

- **Info**: Bold Green
- **Debug**: Bold Yellow
- **Error**: Bold Red
- **Fatal**: Bold Yellow on a Red Background

### Exiting on Fatal Errors

The `Fatal` log level immediately terminates the program after logging.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
