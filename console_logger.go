package logi

import (
	"io"
	"os"
)

// console colors
const (
	Reset = "\033[0m"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Bright colors
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// Backgrounds
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

type ConsoleLogger struct {
	optionss ConsoleLoggerOptions
}

type ConsoleLoggerOptions struct {
	Stdin  io.Reader
	Stdout io.Writer
}

func NewConsoleLogger(option ...ConsoleLoggerOptions) ConsoleLogger {
	if len(option) > 0 && option[0].Stdin != nil && option[0].Stdout != nil {
		return ConsoleLogger{option[0]}
	}

	defaultOptions := ConsoleLoggerOptions{Stdin: os.Stdin, Stdout: os.Stdout}
	return ConsoleLogger{
		optionss: defaultOptions,
	}
}

func prepareEntry(entry Entry) Entry {
	entry.HandleTag(EntryTagLevel, func(e *Entry) string {
		switch e.Level {
		case DebugLevel:
			return Blue + string(entry.Level) + Reset
		case InfoLevel:
			return Green + string(entry.Level) + Reset
		case WarnLevel:
			return BrightYellow + string(entry.Level) + Reset
		case ErrorLevel:
			return Red + string(entry.Level) + Reset
		case FatalLevel:
			return BrightRed + string(entry.Level) + Reset
		}
		return ""
	})
	return entry
}

func (c ConsoleLogger) Info(entry Entry) {
	entry = prepareEntry(entry)
	c.optionss.Stdout.Write([]byte(entry.String()))
}

func (c ConsoleLogger) Debug(entry Entry) {
	entry = prepareEntry(entry)
	c.optionss.Stdout.Write([]byte(entry.String()))
}

func (c ConsoleLogger) Warn(entry Entry) {
	entry = prepareEntry(entry)
	c.optionss.Stdout.Write([]byte(entry.String()))
}

func (c ConsoleLogger) Error(entry Entry) {
	entry = prepareEntry(entry)
	c.optionss.Stdout.Write([]byte(entry.String()))
}

func (c ConsoleLogger) Fatal(entry Entry) {
	entry = prepareEntry(entry)
	c.optionss.Stdout.Write([]byte(entry.String()))
}
