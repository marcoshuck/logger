package logger

import (
	"fmt"
	"io"
	"time"
)

const (
	VerbosityDebug    = 0
	VerbosityInfo     = 1
	VerbosityWarning  = 2
	VerbosityError    = 3
	VerbosityCritical = 4
)

const (
	colorReset  = "\033[0m"
	colorGray   = "\033[37m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
)

type Logger interface {
	Debug(input string)
	Info(input string)
}

type message struct {
	Level int
	Text  string
	Time  time.Time
}

type logger struct {
	Verbosity int
	Writer    io.Writer
}

func (l logger) Debug(input string) {
	l.log(message{
		Level: VerbosityDebug,
		Text:  input,
		Time:  time.Now(),
	})
}

func (l logger) Info(input string) {
	l.log(message{
		Level: VerbosityInfo,
		Text:  input,
		Time:  time.Now(),
	})
}

func (l logger) log(message message) {
	if l.Verbosity > message.Level {
		return
	}
	switch message.Level {
	case VerbosityCritical:
		l.printCritical(message.Time, message.Text)
		break
	case VerbosityError:
		l.printError(message.Time, message.Text)
		break
	case VerbosityWarning:
		l.printWarning(message.Time, message.Text)
		break
	case VerbosityInfo:
		l.printInfo(message.Time, message.Text)
		break
	case VerbosityDebug:
		l.printDebug(message.Time, message.Text)
		break
	default:
		l.printDebug(message.Time, message.Text)
		break
	}
}

func (l logger) printDebug(date time.Time, text string) {
	l.print(date, colorPurple, "DEBUG", text)
}

func (l logger) printInfo(date time.Time, text string) {
	l.print(date, colorBlue, "INFO", text)
}

func (l logger) printWarning(date time.Time, text string) {
	l.print(date, colorYellow, "WARN", text)
}

func (l logger) printError(date time.Time, text string) {
	l.print(date, colorRed, "ERROR", text)
}

func (l logger) printCritical(date time.Time, text string) {
	l.print(date, colorPurple, "CRITICAL", text)
}

func (l logger) print(date time.Time, primaryColor string, key string, text string) {
	fmt.Fprint(l.Writer, colorGray, "[", date.Format(time.RFC822Z), "] ", primaryColor, "[", key, "]", colorReset, " ", text, "\n")
}

func NewLogger(verbosity int, writer io.Writer) Logger {
	return &logger{
		Verbosity: verbosity,
		Writer:    writer,
	}
}
