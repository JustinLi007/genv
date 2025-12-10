package logger

import (
	"fmt"
	"os"
)

type LoggerComponent interface {
	Log(s string)
	Info(s string)
	Warn(s string)
	Error(s string)
}

type LogLevel int

const (
	red        string = "\033[0;31m"
	yellow     string = "\033[0;33m"
	blue       string = "\033[0;34m"
	resetColor string = "\033[0m"
)

type logger struct {
}

func New() LoggerComponent {
	lg := &logger{}
	return lg
}

func (lg *logger) Log(s string) {
	fmt.Fprintf(os.Stderr, "%s", s)
}

func (lg *logger) Info(s string) {
	fmt.Fprintf(os.Stderr, "%sINFO: %s%s", blue, s, resetColor)
}

func (lg *logger) Warn(s string) {
	fmt.Fprintf(os.Stderr, "%sWARN: %s%s", yellow, s, resetColor)
}

func (lg *logger) Error(s string) {
	fmt.Fprintf(os.Stderr, "%sERROR: %s%s", red, s, resetColor)
}

type nullLogger struct {
}

func NewNull() LoggerComponent {
	return &nullLogger{}
}

func (nl *nullLogger) Log(s string) {
}

func (nl *nullLogger) Info(s string) {
}

func (nl *nullLogger) Warn(s string) {
}

func (nl *nullLogger) Error(s string) {
}
