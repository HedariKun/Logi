package logi_test

import (
	"testing"

	"github.com/hedarikun/logi"
)

func TestManager(t *testing.T) {
	LoggerM := logi.New().With("test", []string{"hello", "world"}).With("id", "uuid-3123123")
	LoggerM.AddLogger(logi.NewConsoleLogger())

	fileLogger, err := logi.NewFileLogger(logi.FileLoggerOptions{FolderPath: "./logs", FilePrefix: "test_", FileSuffix: ".log", FileRotateHour: 24, StoreJson: true})
	if err != nil {
		t.Fatal(err)
	}
	LoggerM.AddLogger(fileLogger)

	LoggerM.Info("hello!")
	LoggerM.Debug("hello")
	LoggerM.Warn("hi")
	LoggerM.Error("hii")
	LoggerM.Fatal("VERY HIGH")
}
