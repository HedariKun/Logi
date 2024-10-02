package Logi

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

type IOLogger struct {
	Writer io.Writer
}

func (logger IOLogger) Log(message string) {
	logger.Writer.Write([]byte(message + "\n"))
}

func (logger IOLogger) Error(message string) {
	logger.Writer.Write([]byte(message + "\n"))
}

func (logger IOLogger) Warning(message string) {
	logger.Writer.Write([]byte(message + "\n"))
}

func NewTerminalLogger() Logger {
	logger := IOLogger{}
	logger.Writer = os.Stdout

	return &logger
}

type FileLoggerOption struct {
	Dir      string
	Interval time.Duration
	Prefix   string
}

func NewFileLogger(options FileLoggerOption) (Logger, func()) {
	logger := &IOLogger{}
	var oldfile *os.File
	if _, err := os.Stat(options.Dir); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(options.Dir), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	createLogFile := func() {
		if oldfile != nil {
			oldfile.Close()
		}

		filename := ""

		if options.Prefix != "" {
			filename += options.Prefix + "_"
		}
		filename += "Log_"
		date := time.Now()
		formatedDate := date.Format("2006-01-02 15-04-05:000")
		filename += formatedDate + ".txt"

		file, err := os.Create(options.Dir + filename)
		if err != nil {
			println(err)
		}
		logger.Writer = file
		oldfile = file
	}

	if options.Interval < 5*time.Millisecond {
		options.Interval = 24 * time.Hour
	}

	createLogFile()

	ticker := time.NewTicker(options.Interval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				createLogFile()
			}
		}
	}()

	return logger, ticker.Stop
}
