package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// Infof prints a formatted info message to stdout in green
func Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if viper.GetBool("colorize") {
		msg = color.New(color.FgGreen).Sprint(msg)
	}
	_, _ = fmt.Fprintln(os.Stdout, msg)
}

// Info prints an info message to stdout in green
func Info(msg string) {
	if viper.GetBool("colorize") {
		msg = color.New(color.FgGreen).Sprint(msg)
	}
	_, _ = fmt.Fprintln(os.Stdout, msg)
}

// Error prints an error message to stderr in red
func Error(err error) {
	msg := fmt.Sprintf("%v", err)
	if viper.GetBool("colorize") {
		msg = color.New(color.FgRed).Sprint(msg)
	}
	_, _ = fmt.Fprintf(os.Stderr, "%+v\n", msg)
}

// Fatalf prints a formatted error message to stderr in red and exits
func Fatalf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if viper.GetBool("colorize") {
		msg = color.New(color.FgYellow).Sprint(msg)
	}
	_, _ = fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// Fatal prints an error message to stderr in red and exits
func Fatal(msg string) {
	if viper.GetBool("colorize") {
		msg = color.New(color.FgRed).Sprint(msg)
	}
	_, _ = fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
