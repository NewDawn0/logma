package logma

import (
	"fmt"
	"os"
)

const (
	// Inclusive
	Fatal uint = iota
	Error
	Info
	All
)

// Colours
const info = "\x1b[1;32m"
const dbg = "\x1b[1;33m"
const err = "\x1b[1;31m"
const fatal = "\x1b[1;33;41m"
const nc = "\x1b[0m"

// Allocate it once to save space
var msg string

type Logger struct {
	logLevel uint
}

func NewLogger(lvl uint) Logger {
	return Logger{
		logLevel: lvl,
	}
}
func DefaultLogger() Logger {
	return Logger{
		logLevel: Error,
	}
}
func (self Logger) Dbg(format string, args ...interface{}) {
	if self.logLevel >= All {
		msg = fmt.Sprintf(format, args...)
		fmt.Printf("%s[Dbg]%s %s\n", dbg, nc, msg)
	}
}
func (self Logger) Info(format string, args ...interface{}) {
	if self.logLevel >= Info {
		msg = fmt.Sprintf(format, args...)
		fmt.Printf("%s[Info]%s %s\n", info, nc, msg)
	}
}
func (self Logger) Error(format string, args ...interface{}) {
	if self.logLevel >= Error {
		msg = fmt.Sprintf(format, args...)
		fmt.Printf("%s[Error]%s %s\n", err, nc, msg)
	}
}
func (self Logger) Fatal(format string, args ...interface{}) {
	if self.logLevel >= Fatal {
		msg = fmt.Sprintf(format, args...)
		fmt.Printf("%s[Fatal]%s %s%s%s\n", fatal, nc, err, msg, nc)
		os.Exit(1)
	}
}
