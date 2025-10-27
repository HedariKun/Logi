package logi

import (
	"os"
	"strings"
	"time"
)

type FileLogger struct {
	options        FileLoggerOptions
	root           *os.Root
	openFile       *os.File
	currentLogTime time.Time
}

type FileLoggerOptions struct {
	FolderPath     string
	FileRotateHour int
	FileSuffix     string
	FilePrefix     string
	StoreJson      bool
}

func NewFileLogger(options FileLoggerOptions) (*FileLogger, error) {
	logger := FileLogger{options: options}

	if options.FileRotateHour <= 0 {
		options.FileRotateHour = 24
	}

	if options.FilePrefix == "" {
		options.FilePrefix = "log_"
	}

	if options.FileSuffix == "" {
		options.FileSuffix = ".log"
	}
	// TODO: make all error handling self contained, rn let the user handle all the errors
	err := os.MkdirAll(options.FolderPath, 0755)

	if err != nil {
		return nil, err
	}

	// TODO: handle proper root closing
	logger.root, err = os.OpenRoot(options.FolderPath)

	if err != nil {
		return nil, err
	}
	err = fileCycle(&logger)
	if err != nil {
		return nil, err
	}

	return &logger, nil
}

func (f *FileLogger) Close() error {
	return f.root.Close()
}

func (f *FileLogger) handleEntry(entry Entry) {
	if f.options.StoreJson {
		data, err := entry.Json()
		// TODO: let all error handling being centralized by Logi
		if err != nil {
			return
		}

		f.openFile.Write([]byte(data + "\n"))
	} else {
		f.openFile.Write([]byte(entry.String()))
	}
}

func (f *FileLogger) Info(entry Entry) {
	f.handleEntry(entry)
}

func (f *FileLogger) Debug(entry Entry) {
	f.handleEntry(entry)
}

func (f *FileLogger) Warn(entry Entry) {
	f.handleEntry(entry)
}

func (f *FileLogger) Error(entry Entry) {
	f.handleEntry(entry)
}

func (f *FileLogger) Fatal(entry Entry) {
	f.handleEntry(entry)
}

func fileCycle(logger *FileLogger) error {
	files, err := os.ReadDir(logger.options.FolderPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileNameTime := file.Name()
		fileNameTime = strings.Replace(fileNameTime, logger.options.FilePrefix, "", 1)
		fileNameTime = strings.Replace(fileNameTime, logger.options.FileSuffix, "", 1)

		fileTime, err := time.Parse("2006-01-02T15-04-05Z0700", fileNameTime)
		if err != nil {
			return err
		}
		// TODO: let the user decide if it is hour, minute or second from options
		if time.Since(fileTime) < time.Hour*time.Duration(logger.options.FileRotateHour) {
			println("hello :>", time.Since(fileTime), time.Now().Format("2006-01-02T15-04-05Z0700"), fileTime.Format("2006-01-02T15-04-05Z0700"), time.Now().UnixMilli(), fileTime.UnixMilli(), time.Hour*time.Duration(logger.options.FileRotateHour), logger.options.FileRotateHour, time.Duration(logger.options.FileRotateHour), time.Hour)
			logger.currentLogTime = fileTime
			logger.openFile, err = os.OpenFile(logger.options.FolderPath+"/"+file.Name(), os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			break
		}

	}
	// TODO: edit this as well
	if logger.openFile == nil || (logger.openFile != nil && time.Since(logger.currentLogTime) >= time.Hour*time.Duration(logger.options.FileRotateHour)) {
		if logger.openFile != nil {
			logger.openFile.Close()
			logger.openFile = nil
		}
		file, err := logger.root.Create(logger.options.FilePrefix + time.Now().Format("2006-01-02T15-04-05Z0700") + logger.options.FileSuffix)
		if err != nil {
			return err
		}
		logger.openFile = file
		logger.currentLogTime = time.Now()
	}

	return nil
}
